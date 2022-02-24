package uadmin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type initialDataRecords []map[string]interface{}

type initialData struct {
	Init   []string                      `json:"init"`
	Data   map[string]initialDataRecords `json:"data"`
	Finish []string                      `json:"finish"`
}

// loadInitialData reads a file named initial_data.json and
// saves its content in the database
func loadInitialData() error {
	buf, err := ioutil.ReadFile("initial_data.json")
	if err != nil {
		return nil
	}

	// Load json daa into struct
	data := initialData{}
	err = json.Unmarshal(buf, &data)
	if err != nil {
		return fmt.Errorf("loadInitialData.Unmarshal: Error parsing json file. %s", err)
	}

	// Execute SQL in Init section
	for _, SQL := range data.Init {
		err = db.Exec(SQL).Error
		if err != nil {
			return fmt.Errorf("loadInitialData.Exec: Error in in Init section (%s). %s", SQL, err)
		}
	}

	// Load data
	for table, records := range data.Data {
		// get modelname from table name
		// For the record:
		//   - Name       :  OrderItem
		//   - DisplayName:  Order Items
		//   - ModelName  :  orderitem
		//   - TableName  :  order_items
		tabeFound := false
		for k, v := range Schema {
			// check if table is a ModelName
			if table == k {
				tabeFound = true
				break
			}
			// check if the table is a database TableName
			// and convert it into modelname
			if v.TableName == table {
				table = k
				tabeFound = true
				break
			}
			// check if the table is a Name
			// and convert it into modelname
			if v.Name == table {
				table = k
				tabeFound = true
				break
			}
		}
		if !tabeFound {
			return fmt.Errorf("loadInitialData: Table not found for (%s)", table)
		}

		// Put records into Model Array
		modelArray, _ := NewModelArray(table, true)
		buf, _ = json.Marshal(records)
		err = json.Unmarshal(buf, modelArray.Interface())
		if err != nil {
			return fmt.Errorf("loadInitialData.Unmarshal: Error parsing Data records in (%s). %s", table, err)
		}

		// Save records
		for i := 0; i < modelArray.Elem().Len(); i++ {
			err = Save(modelArray.Elem().Index(i).Addr().Interface())
			if err != nil {
				return fmt.Errorf("loadInitialData.Save: Error in %s[%d]. %s", table, i, err)
			}
		}
	}

	// Execute SQL in Finish section
	for _, SQL := range data.Finish {
		err = db.Exec(SQL).Error
		if err != nil {
			return fmt.Errorf("loadInitialData.Exec: Error in in Finish section (%s). %s", SQL, err)
		}
	}

	return nil
}
