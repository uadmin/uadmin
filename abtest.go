package uadmin

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type ABTestType int

func (ABTestType) Static() ABTestType {
	return 1
}

func (ABTestType) Model() ABTestType {
	return 2
}

// ModelList a list of registered models
type ModelList int

// FieldList is a list of fields from schema for a registered model
type FieldList int

var staticABTests map[string][]struct {
	v     string
	vid   uint
	imp   uint
	click uint
	group string
}

var modelABTests map[string][]struct {
	v     string
	vid   uint
	fname int
	pk    uint
	imp   uint
	click uint
	group string
}

var abTestsMutex = sync.Mutex{}
var abTestCount = 0

// ABTest is a model that stores an A/B test
type ABTest struct {
	Model
	Name        string     `uadmin:"required"`
	Type        ABTestType `uadmin:"required"`
	StaticPath  string
	ModelName   ModelList
	Field       FieldList
	PrimaryKey  int
	Active      bool
	Group       string
	ResetABTest string `uadmin:"link"`
}

// Save !
func (a *ABTest) Save() {
	if a.ResetABTest == "" {
		a.ResetABTest = RootURL + "api/d/abtest/method/Reset/" + fmt.Sprint(a.ID) + "/?$next=$back"
	}
	Save(a)
	abTestCount = Count([]ABTest{}, "active = ?", true)
}

func loadModels(a interface{}, u *User) []Choice {
	c := []Choice{}
	for i, m := range modelList {
		c = append(c, Choice{K: uint(i), V: getModelName(m)})
	}
	return c
}

func loadFields(a interface{}, u *User) []Choice {
	m, ok := a.(ABTest)
	if !ok {
		mp, ok := a.(*ABTest)
		if !ok {
			Trail(ERROR, "loadFields Unable to cast a to ABTest")
			return []Choice{}
		}
		m = *mp
	}

	if m.Type != m.Type.Model() {
		return []Choice{}
	}

	s := Schema[getModelName(modelList[int(m.ModelName)])]
	c := []Choice{}
	for i, f := range s.Fields {
		c = append(c, Choice{K: uint(i), V: f.Name})
	}
	return c
}

func syncABTests() {
	// Check if there are stats to save to the DB
	abTestsMutex.Lock()
	if staticABTests != nil {
		tx := db.Begin()
		for _, v := range staticABTests {
			for i := range v {
				if v[i].imp != 0 || v[i].click != 0 {
					// store results to DB
					tx.Exec("UPDATE ab_test_values SET impressions = impressions + ?, clicks = clicks + ? WHERE id = ?", v[i].imp, v[i].click, v[i].vid)
				}
			}
		}

		for _, v := range modelABTests {
			for i := range v {
				if v[i].imp != 0 || v[i].click != 0 {
					// store results to DB
					tx.Exec("UPDATE ab_test_values SET impressions = impressions + ?, clicks = clicks + ? WHERE id = ?", v[i].imp, v[i].click, v[i].vid)
				}
			}
		}
		tx.Commit()
	}
	staticABTests = map[string][]struct {
		v     string
		vid   uint
		imp   uint
		click uint
		group string
	}{}

	modelABTests = map[string][]struct {
		v     string
		vid   uint
		fname int
		pk    uint
		imp   uint
		click uint
		group string
	}{}

	tests := []ABTest{}
	Filter(&tests, "active = ?", true)

	// Process Static AB Tests
	for _, t := range tests {
		if t.Type != t.Type.Static() {
			continue
		}
		values := []ABTestValue{}
		Filter(&values, "ab_test_id = ? AND active = ?", t.ID, true)
		tempList := []struct {
			v     string
			vid   uint
			imp   uint
			click uint
			group string
		}{}
		for _, v := range values {
			tempList = append(tempList, struct {
				v     string
				vid   uint
				imp   uint
				click uint
				group string
			}{v: v.Value, vid: v.ID, group: t.Group})
		}
		staticABTests[t.StaticPath] = tempList
	}

	// Process Models AB Tests
	for _, t := range tests {
		if t.Type != t.Type.Model() {
			continue
		}
		schema := Schema[getModelName(modelList[int(t.ModelName)])]
		fName := schema.Fields[int(t.Field)].Name
		values := []ABTestValue{}
		Filter(&values, "ab_test_id = ? AND active = ?", t.ID, true)
		tempList := []struct {
			v     string
			vid   uint
			fname int
			pk    uint
			imp   uint
			click uint
			group string
		}{}
		for _, v := range values {
			tempList = append(tempList, struct {
				v     string
				vid   uint
				fname int
				pk    uint
				imp   uint
				click uint
				group string
			}{v: v.Value, vid: v.ID, group: t.Group, pk: uint(t.PrimaryKey), fname: int(t.Field)})
		}
		modelABTests[schema.ModelName+"__"+fName+"__"+fmt.Sprint(t.PrimaryKey)] = tempList
	}
	abTestsMutex.Unlock()
}

// ABTestClick is a function to register a click for an ABTest group
func ABTestClick(r *http.Request, group string) {
	go func() {
		abt := getABT(r)
		var index int
		abTestsMutex.Lock()
		for k, v := range staticABTests {
			if len(v) != 0 && v[0].group == group {
				index = abt % len(v)
				v[index].click++
				staticABTests[k] = v
			}
		}
		for k, v := range modelABTests {
			if len(v) != 0 && v[0].group == group {
				index = abt % len(v)
				v[index].click++
				modelABTests[k] = v
			}
		}
		abTestsMutex.Unlock()
	}()
}

func getABT(r *http.Request) int {
	c, err := r.Cookie("abt")
	if err != nil || c == nil {
		now := time.Now().AddDate(0, 0, 1)
		/*http.SetCookie(&http.Cookie{
			Name:    "abt",
			Value:   fmt.Sprint(now.Second()),
			Path:    "/",
			Expires: time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()),
		})
		*/
		return now.Second()
	}

	v, _ := strconv.ParseInt(c.Value, 10, 64)
	return int(v)
}

// Reset resets the impressions and clicks to 0 based on a specified
// AB Test ID
func (a ABTest) Reset() {
	abTestsMutex.Lock()
	abtestValue := ABTestValue{}
	Update(&abtestValue, "Impressions", 0, "ab_test_id = ?", a.ID)
	Update(&abtestValue, "Clicks", 0, "ab_test_id = ?", a.ID)
	abTestsMutex.Unlock()
}
