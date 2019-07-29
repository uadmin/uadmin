package uadmin

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
	"time"
)

// Approval is a model that stores approval data
type Approval struct {
	Model
	ModelName           string `uadmin:"read_only"`
	ModelPK             uint   `uadmin:"read_only"`
	ColumnName          string `uadmin:"read_only"`
	OldValue            string `uadmin:"read_only"`
	NewValue            string
	NewValueDescription string    `uadmin:"read_only"`
	ChangedBy           string    `uadmin:"read_only"`
	ChangeDate          time.Time `uadmin:"read_only"`
	ApprovalAction      ApprovalAction
	ApprovalBy          string     `uadmin:"read_only"`
	ApprovalDate        *time.Time `uadmin:"read_only"`
	ViewRecord          string     `uadmin:"link"`
	UpdatedBy           string     `uadmin:"read_only;hidden;list_exclude"`
}

func (a *Approval) String() string {
	return fmt.Sprintf("%s.%s %d", a.ModelName, a.ColumnName, a.ModelPK)
}

// Save overides save
func (a *Approval) Save() {
	if a.ViewRecord == "" {
		a.ViewRecord = RootURL + a.ModelName + "/" + fmt.Sprint(a.ModelPK)
	}
	if Schema[a.ModelName].FieldByName(a.ColumnName).Type == cLIST {
		m, _ := NewModel(a.ModelName, false)
		intVal, _ := strconv.ParseInt(a.NewValue, 10, 64)
		m.FieldByName(a.ColumnName).SetInt((intVal))
		a.NewValueDescription = GetString(m.FieldByName(a.ColumnName).Interface())
	} else if Schema[a.ModelName].FieldByName(a.ColumnName).Type == cFK {
		m, _ := NewModel(strings.ToLower(Schema[a.ModelName].FieldByName(a.ColumnName).TypeName), true)
		Get(m.Interface(), "id = ?", a.NewValue)
		a.NewValueDescription = GetString(m.Interface())
	} else {
		a.NewValueDescription = a.NewValue
	}

	// Run Approval handle func
	saveApproval := true
	if ApprovalHandleFunc != nil {
		saveApproval = ApprovalHandleFunc(a)
	}

	// Process approval based on the action
	old := Approval{}
	if a.ID != 0 {
		Get(&old, "id = ?", a.ID)
	}
	if old.ApprovalAction != a.ApprovalAction {
		a.ApprovalBy = a.UpdatedBy
		now := time.Now()
		a.ApprovalDate = &now
		m, _ := NewModelArray(a.ModelName, true)
		model, _ := NewModel(a.ModelName, false)
		if a.ApprovalAction == a.ApprovalAction.Approved() {
			if model.FieldByName(a.ColumnName).Type().String() == "*time.Time" && a.NewValue == "" {
				Update(m.Interface(), gorm.ToColumnName(a.ColumnName), nil, "id = ?", a.ModelPK)
			} else if Schema[a.ModelName].FieldByName(a.ColumnName).Type == cFK {
				Update(m.Interface(), gorm.ToColumnName(a.ColumnName)+"_id", a.NewValue, "id = ?", a.ModelPK)
			} else {
				Update(m.Interface(), gorm.ToColumnName(a.ColumnName), a.NewValue, "id = ?", a.ModelPK)
			}
		} else {
			if model.FieldByName(a.ColumnName).Type().String() == "*time.Time" && a.OldValue == "" {
				Update(m.Interface(), gorm.ToColumnName(a.ColumnName), nil, "id = ?", a.ModelPK)
			} else if Schema[a.ModelName].FieldByName(a.ColumnName).Type == cFK {
				Update(m.Interface(), gorm.ToColumnName(a.ColumnName)+"_id", a.OldValue, "id = ?", a.ModelPK)
			} else {
				Update(m.Interface(), gorm.ToColumnName(a.ColumnName), a.OldValue, "id = ?", a.ModelPK)
			}
		}
	}

	if !saveApproval {
		return
	}

	Save(a)
}
