package uadmin

import (
	//"io"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
	//"path/filepath"
	"strconv"
	"sync"
)

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
	Filter(&tests, "active = ?", TestType(0).Static(), true)

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

func ABTestClick(r *http.Request, group string) {
	go func() {
		abt := getABT(r)
		var index int
		for k, v := range staticABTests {
			if len(v) != 0 && v[0].group == group {
				abTestsMutex.Lock()
				index = abt % len(v)
				v[index].click++
				staticABTests[k] = v
				abTestsMutex.Unlock()
			}
		}
	}()
}

func getABT(r *http.Request) int {
	c, err := r.Cookie("abt")
	if err != nil || c == nil {
		Trail(DEBUG, "ERROR:%s", err)
		return 0
	}

	v, _ := strconv.ParseInt(c.Value, 10, 64)
	return int(v)
}

func containsDotDot(v string) bool {
	if !strings.Contains(v, "..") {
		return false
	}
	for _, ent := range strings.FieldsFunc(v, func(r rune) bool { return r == '/' || r == '\\' }) {
		if ent == ".." {
			return true
		}
	}
	return false
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	if containsDotDot(r.URL.Path) {
		w.WriteHeader(404)
		return
	}
	var modTime time.Time
	ab := false
	for k := range staticABTests {
		if k == r.URL.Path && len(staticABTests[k]) != 0 {
			index := getABT(r) % len(staticABTests[k])
			r.URL.Path = staticABTests[k][index].v
			// TODO: Change max-age to midnight
			w.Header().Add("Cache-Control", "private, max-age=1")
			modTime = time.Now()
			ab = true

			go func() {
				abTestsMutex.Lock()
				t := staticABTests[k]
				t[index].imp++
				staticABTests[k] = t
				abTestsMutex.Unlock()
			}()
			break
		}
	}

	f, err := os.Open("." + r.URL.Path)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	if !ab {
		stat, err := os.Stat("." + r.URL.Path)
		if err != nil || stat.IsDir() {
			w.WriteHeader(404)
			return
		}
		modTime = stat.ModTime()
	}

	http.ServeContent(w, r, "."+r.URL.Path, modTime, f)
}
