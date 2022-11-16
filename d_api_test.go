package uadmin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"time"
)

// TestdAPI to test dAPI
func (t *UAdminTests) TestDAPI() {
	u1 := &User{
		Username:     "u1",
		Password:     "u1",
		Active:       true,
		RemoteAccess: true,
		Admin:        true,
	}
	u1.Save()
	s1 := &Session{
		Active:    true,
		UserID:    u1.ID,
		LoginTime: time.Now(),
	}
	s1.GenerateKey()
	s1.Save()

	e := []struct {
		url      string
		session  string
		validate func(string) string
	}{
		{
			"/api/d",
			s1.Key,
			func(v string) string {
				if v != dAPIHelp {
					return fmt.Sprintf("Invalid return for dAPI url=%%s")
				}
				return ""
			},
		},
		{
			"/api/d/$allmodels",
			s1.Key,
			func(v string) string {
				obj := map[string]interface{}{}
				json.Unmarshal([]byte(v), &obj)
				if result, ok := obj["result"].([]interface{}); !ok {
					return fmt.Sprintf("Invalid return for dAPI url=%%s. No 'result' in response")
				} else if len(result) != 17 {
					return fmt.Sprintf("Invalid length of 'result' dAPI url=%%s. Expected %d got %d", 17, len(result))
				}
				return ""
			},
		},
		{
			"/api/d/user/read",
			s1.Key,
			func(v string) string {
				obj := map[string]interface{}{}
				json.Unmarshal([]byte(v), &obj)
				if result, ok := obj["result"].([]interface{}); !ok {
					return fmt.Sprintf("Invalid return for dAPI url=%%s. No 'result' in response")
				} else if len(result) != 2 {
					return fmt.Sprintf("Invalid length of 'result' dAPI url=%%s. Expected %d got %d", 2, len(result))
				}
				return ""
			},
		},
		{
			"/api/d/user/read/" + fmt.Sprint(u1.ID),
			s1.Key,
			func(v string) string {
				obj := map[string]interface{}{}
				json.Unmarshal([]byte(v), &obj)
				if result, ok := obj["result"]; !ok {
					return fmt.Sprintf("Invalid return for dAPI url=%%s. No 'result' in response")
				} else if _, ok := result.(map[string]interface{}); !ok {
					return fmt.Sprintf("Invalid value 'result' dAPI url=%%s. Expected map got %v", result)
				}
				return ""
			},
		},
		{
			"/api/d/user/none_existing_command",
			s1.Key,
			func(v string) string {
				obj := map[string]interface{}{}
				json.Unmarshal([]byte(v), &obj)
				if result, ok := obj["status"].(string); !ok {
					return fmt.Sprintf("Invalid return for dAPI url=%%s. No 'status' in response")
				} else if result != "error" {
					return fmt.Sprintf("Invalid value of 'status' dAPI url=%%s. Expected %s got %s", "error", result)
				}
				return ""
			},
		},
		{
			"/api/d/none_existing_model/read",
			s1.Key,
			func(v string) string {
				obj := map[string]interface{}{}
				json.Unmarshal([]byte(v), &obj)
				if result, ok := obj["status"].(string); !ok {
					return fmt.Sprintf("Invalid return for dAPI url=%%s. No 'status' in response")
				} else if result != "error" {
					return fmt.Sprintf("Invalid value of 'status' dAPI url=%%s. Expected %s got %s", "error", result)
				}
				return ""
			},
		},
		{
			"/api/d/testmodela/add?_name=test_dAPI&x-csrf-token=" + s1.Key,
			s1.Key,
			func(v string) string {
				obj := map[string]interface{}{}
				json.Unmarshal([]byte(v), &obj)
				if result, ok := obj["status"].(string); !ok {
					return fmt.Sprintf("Invalid return for dAPI url=%%s. No 'status' in response")
				} else if result != "ok" {
					return fmt.Sprintf("Invalid value of 'status' dAPI url=%%s. Expected %s got %s. %s", "ok", result, v)
				}
				if result, ok := obj["rows_count"]; !ok {
					return fmt.Sprintf("Invalid return for dAPI url=%%s. No 'rows_count' in response")
				} else if result.(float64) != 1 {
					return fmt.Sprintf("Invalid value of 'rows_count' dAPI url=%%s. Expected %f got %f", 1.0, result.(float64))
				}
				if result, ok := obj["id"]; !ok {
					return fmt.Sprintf("Invalid return for dAPI url=%%s. No 'id' in response")
				} else if len(result.([]interface{})) != 1 {
					return fmt.Sprintf("Invalid length of 'id' dAPI url=%%s. Expected %d got %d", 1, len(result.([]interface{})))
				}
				return ""
			},
		},
		{
			"/api/d/testmodela/edit?name=test_dAPI&_name=test_dAPI2&x-csrf-token=" + s1.Key,
			s1.Key,
			func(v string) string {
				obj := map[string]interface{}{}
				json.Unmarshal([]byte(v), &obj)
				if result, ok := obj["status"].(string); !ok {
					return fmt.Sprintf("Invalid return for dAPI url=%%s. No 'status' in response")
				} else if result != "ok" {
					return fmt.Sprintf("Invalid value of 'status' dAPI url=%%s. Expected %s got %s", "ok", result)
				}
				return ""
			},
		},
		{
			"/api/d/testmodela/delete?name=test_dAPI2&x-csrf-token=" + s1.Key,
			s1.Key,
			func(v string) string {
				obj := map[string]interface{}{}
				json.Unmarshal([]byte(v), &obj)
				if result, ok := obj["status"].(string); !ok {
					return fmt.Sprintf("Invalid return for dAPI url=%%s. No 'status' in response")
				} else if result != "ok" {
					return fmt.Sprintf("Invalid value of 'status' dAPI url=%%s. Expected %s got %s", "ok", result)
				}
				return ""
			},
		},
		{
			"/api/d/testmodela/schema",
			s1.Key,
			func(v string) string {
				obj := map[string]interface{}{}
				json.Unmarshal([]byte(v), &obj)
				if result, ok := obj["status"].(string); !ok {
					return fmt.Sprintf("Invalid return for dAPI url=%%s. No 'status' in response")
				} else if result != "ok" {
					return fmt.Sprintf("Invalid value of 'status' dAPI url=%%s. Expected %s got %s", "ok", result)
				}
				if result, ok := obj["result"]; !ok {
					return fmt.Sprintf("Invalid return for dAPI url=%%s. No 'result' in response")
				} else if _, ok := result.(map[string]interface{}); !ok {
					return fmt.Sprintf("Invalid value 'result' dAPI url=%%s. Expected map got %v", result)
				}
				return ""
			},
		},
	}

	for i := range e {
		r := httptest.NewRequest("GET", e[i].url, nil)

		if e[i].session != "" {
			c := http.Cookie{}
			c.Name = "session"
			c.Value = e[i].session
			r.AddCookie(&c)
		}

		w := httptest.NewRecorder()

		apiHandler(w, r)

		buf, err := ioutil.ReadAll(w.Result().Body)
		if err != nil {
			t.Errorf("Unable to read dAPI response for example %d", i)
			continue
		}

		if msg := e[i].validate(string(buf)); msg != "" {
			t.Errorf(msg+" in example %d", e[i].url, i)
		}
	}
	Delete(s1)
	Delete(u1)
}
