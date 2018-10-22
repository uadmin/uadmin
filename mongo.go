package uadmin

import (
	"fmt"
	"log"
	"strings"

	"reflect"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoSettings !
type MongoSettings struct {
	Name  string
	IP    string
	Debug bool
}

// MongoDB !
var MongoDB *MongoSettings

var mongoSession *mgo.Session

// MongoModel !
type MongoModel struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
}

func (m MongoModel) getSession() (*mgo.Session, error) {
	session, err := mgo.Dial(MongoDB.IP)
	if err != nil {
		log.Println("ERROR: ds.Model.getSession", err)
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	return session, nil
}

// GetCol !
func (m MongoModel) GetCol(a interface{}, ColNameExtra string) (*mgo.Collection, error) {
	var err error
	if mongoSession == nil {
		mongoSession, err = m.getSession()
		if err != nil {
			return nil, err
		}
	}
	colName := m.getColName(a) + ColNameExtra
	c := mongoSession.DB(MongoDB.Name).C(colName)
	return c, nil
}

func (m MongoModel) getColName(a interface{}) string {
	colName := reflect.TypeOf(a).String()
	if strings.Contains(colName, ".") {
		arrColName := strings.Split(colName, ".")
		colName = arrColName[len(arrColName)-1]
	}
	colName = strings.TrimPrefix(colName, "*")
	colName = strings.TrimPrefix(colName, "[]")
	colName = strings.TrimPrefix(colName, "*")
	colName = strings.ToLower(colName)
	if MongoDB.Debug {
		fmt.Println("colName", colName)
	}
	return colName
}

// Save !
func (m MongoModel) Save(a interface{}, ColNameExtra string) {
	var err error
	c, _ := m.GetCol(a, ColNameExtra)

	if m.ID == "" {
		if MongoDB.Debug {
			fmt.Println("Inserting", a)
		}
		err = c.Insert(a)
		if err != nil {
			log.Println("ERROR: ds.Model.Save INSERT", err, a)
		}
	} else {
		err = c.UpdateId(m.ID, a)
		if MongoDB.Debug {
			fmt.Println("Updating", a)
		}
		if err != nil {
			log.Println("ERROR: ds.Model.Save UPDATE", err)
		}
	}
}

// Get !
func (m MongoModel) Get(filter interface{}, a interface{}, ColNameExtra string) error {
	c, err := m.GetCol(a, ColNameExtra)
	if err != nil {
		return err
	}

	err = c.Find(filter).One(a)
	return err
}

// Filter !
func (m MongoModel) Filter(filter interface{}, a interface{}, ColNameExtra string) error {
	c, err := m.GetCol(a, ColNameExtra)
	if err != nil {
		return err
	}

	err = c.Find(filter).All(a)
	return err
}

// Delete !
func (m MongoModel) Delete(a interface{}, ColNameExtra string) error {
	c, err := m.GetCol(a, ColNameExtra)
	if err != nil {
		return err
	}

	if MongoDB.Debug {
		fmt.Println("DELETE FROM", c.Name, m.ID)
	}

	err = c.RemoveId(m.ID)
	return err
}

// All !
func (m MongoModel) All(a interface{}, ColNameExtra string) error {
	return m.Filter(bson.M{}, a, ColNameExtra)
}

// Count !
func (m MongoModel) Count(filter interface{}, a interface{}, ColNameExtra string) int {
	c, err := m.GetCol(a, ColNameExtra)
	if err != nil {
		return -1
	}

	count, err := c.Find(filter).Count()
	if err != nil {
		return -1
	}
	return count
}

// Query !
func (m MongoModel) Query(filter interface{}, a interface{}, ColNameExtra string) *mgo.Query {
	c, err := m.GetCol(a, ColNameExtra)
	if err != nil {
		return nil
	}

	q := c.Find(filter)
	if err != nil {
		return nil
	}
	return q
}
