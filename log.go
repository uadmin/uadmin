package uadmin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"
)

// Action !
type Action int

func (a Action) Read() Action {
	return 1
}

// Added @
func (a Action) Added() Action {
	return 2
}

// Modified !
func (a Action) Modified() Action {
	return 3
}

// Deleted !
func (a Action) Deleted() Action {
	return 4
}

// LoginSuccessful !
func (a Action) LoginSuccessful() Action {
	return 5
}

// LoginDenied !
func (a Action) LoginDenied() Action {
	return 6
}

// Logout !
func (a Action) Logout() Action {
	return 7
}

// PasswordResetRequest !
func (a Action) PasswordResetRequest() Action {
	return 8
}

// PasswordResetDenied !
func (a Action) PasswordResetDenied() Action {
	return 9
}

// PasswordResetSuccessful !
func (a Action) PasswordResetSuccessful() Action {
	return 10
}

// GetSchema !
func (a Action) GetSchema() Action {
	return 11
}

// Custom !
func (a Action) Custom() Action {
	return 99
}

// Log !
type Log struct {
	Model
	Username  string    `uadmin:"filter;read_only"`
	Action    Action    `uadmin:"filter;read_only"`
	TableName string    `uadmin:"filter;read_only"`
	TableID   int       `uadmin:"filter;read_only"`
	Activity  string    `uadmin:"code;read_only" gorm:"type:longtext"`
	RollBack  string    `uadmin:"link;"`
	CreatedAt time.Time `uadmin:"filter;read_only"`
}

func (l Log) String() string {
	return fmt.Sprint(l.ID)
}

// Save !
func (l *Log) Save() {
	Save(l)
	if l.Action == l.Action.Modified() || l.Action == l.Action.Deleted() {
		l.RollBack = RootURL + "revertHandler/?log_id=" + fmt.Sprint(l.ID)
	}
	Save(l)
}

// ParseRecord !
func (l *Log) ParseRecord(a reflect.Value, modelName string, ID uint, user *User, action Action, r *http.Request) (err error) {
	modelName = strings.ToLower(modelName)
	s, ok := getSchema(modelName)
	if !ok {
		errMsg := fmt.Sprintf("Unable to find schema (%s)", modelName)
		Trail(ERROR, errMsg)
		return fmt.Errorf(errMsg)
	}
	l.Username = user.Username
	l.TableName = modelName
	l.TableID = int(ID)
	l.Action = action

	// Check if the value passed is a pointer
	if a.Kind() == reflect.Ptr {
		a = a.Elem()
	}

	jsonifyValue := map[string]string{
		"_IP": r.RemoteAddr,
	}
	for _, f := range s.Fields {
		if !f.IsMethod {
			if f.Type == cFK {
				jsonifyValue[f.Name+"ID"] = fmt.Sprint(a.FieldByName(f.Name + "ID").Interface())
			} else if f.Type == cDATE {
				val := time.Time{}
				if a.FieldByName(f.Name).Type().Kind() == reflect.Ptr {
					if a.FieldByName(f.Name).IsNil() {
						jsonifyValue[f.Name] = ""
					} else {
						val, _ = a.FieldByName(f.Name).Elem().Interface().(time.Time)
						jsonifyValue[f.Name] = val.Format("2006-01-02 15:04:05 -0700")
					}

				} else {
					val, _ = a.FieldByName(f.Name).Interface().(time.Time)
					jsonifyValue[f.Name] = val.Format("2006-01-02 15:04:05 -0700")
				}

			} else {
				jsonifyValue[f.Name] = fmt.Sprint(a.FieldByName(f.Name).Interface())
			}

		}
	}
	json, _ := json.Marshal(jsonifyValue)
	l.Activity = string(json)

	return nil
}

// SignIn !
func (l *Log) SignIn(user string, action Action, r *http.Request) (err error) {

	l.Username = user
	l.Action = action
	jsonifyValue := map[string]string{
		"IP":           r.RemoteAddr,
		"Login-Status": r.FormValue("login-status"),
	}
	for k, v := range r.Header {
		jsonifyValue[k] = strings.Join(v, ";")
	}

	json, _ := json.Marshal(jsonifyValue)
	l.Activity = string(json)

	return nil
}

// PasswordReset !
func (l *Log) PasswordReset(user string, action Action, r *http.Request) (err error) {

	l.Username = user
	l.Action = action
	jsonifyValue := map[string]string{
		"IP":           r.RemoteAddr,
		"Reset-Status": r.FormValue("reset-status"),
	}
	for k, v := range r.Header {
		jsonifyValue[k] = strings.Join(v, ";")
	}

	json, _ := json.Marshal(jsonifyValue)
	l.Activity = string(json)

	return nil
}
