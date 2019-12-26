package uadmin

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

// getSchema returns a schema of a form
func getSchema(a interface{}) (s ModelSchema, ok bool) {
	// Get type of the models
	t := reflect.TypeOf(a)

	modelName := getModelName(a)
	if t.Kind() == reflect.String {
		modelName = strings.ToLower(a.(string))
	}

	// Check if the models has been processed and return it from global schema
	if val, ok := Schema[modelName]; ok {
		cpy, ok := deepCopy(val).(ModelSchema)
		return cpy, ok
	}

	if t.Kind() != reflect.Struct {
		Trail(WARNING, string(debug.Stack()))
		Trail(WARNING, "Unable to get schema for none struct type (%s). %v", t.Name(), a)
		return
	}

	// Get basic information about the model
	s.Name = t.Name()
	s.ModelName = strings.ToLower(t.Name())
	s.DisplayName = getDisplayName(t.Name())
	s.TableName = GetDB().NewScope(a).TableName()

	// Analize the fields of the model and add them to the fields list
	s.Fields = []F{}

	// Add inlines to schema
	// Make a list of schema inline
	s.Inlines = []*ModelSchema{}

	// Get each inline and add it to the list of inlines
	for _, i := range inlines[s.ModelName] {
		inlineSchema, _ := getSchema(i)
		s.Inlines = append(s.Inlines, &inlineSchema)
	}

	now := time.Now()
	// Define generic types of fields
	SType := reflect.TypeOf("")          // String
	DType := reflect.TypeOf(now)         // Date
	DType1 := reflect.TypeOf(&now)       //
	BType := reflect.TypeOf(true)        // Bool
	NType := reflect.TypeOf(int(0))      // Number
	NType1 := reflect.TypeOf(int32(0))   //
	NType2 := reflect.TypeOf(int64(0))   //
	NType3 := reflect.TypeOf(uint(0))    //
	NType4 := reflect.TypeOf(uint32(0))  //
	NType5 := reflect.TypeOf(uint64(0))  //
	NType6 := reflect.TypeOf(float32(0)) //
	NType7 := reflect.TypeOf(float64(0)) //

	// Loop over the fields of the model and get schema information
	for index := 0; index < t.NumField(); index++ {
		// If the field is private, skip it
		if strings.ToLower(string(t.Field(index).Name[0])) == string(t.Field(index).Name[0]) {
			continue
		}

		// Initialize the field
		f := F{
			Translations: []translation{},
		}

		// Get field's meta data
		f.Name = t.Field(index).Name
		f.DisplayName = getDisplayName(t.Field(index).Name)
		f.ColumnName = gorm.ToColumnName(t.Field(index).Name)

		// Get uadmin tag from the field
		tagList := strings.Split(t.Field(index).Tag.Get("uadmin"), ";")
		tagMap := map[string]string{}
		tagParts := map[string]string{}

		// Rejoin items in the tagList if the semi colon was escaped
		// For example: `uadmin:"help:use ; to separate your sales"`
		// This will not work, you should escape the semi colon like this:
		// `uadmin:"help:use \; to separate your sales"`
		if len(tagList) > 0 {
			//tagMap = map[string]string{}
			var skipNext bool
			var i int
			for i = range tagList {
				skipNext = false
				tagParts[fmt.Sprint(i)] = tagList[i]
				// Check if the escape is not the end of the uadmin tag and it is not escaped
				if strings.HasSuffix(tagList[i], "\\") && i < len(tagList)-1 && !strings.HasSuffix(tagList[i], "\\\\") {
					tagParts[fmt.Sprint(i)] += tagList[i+1]
					skipNext = true
				}
				tagParts[fmt.Sprint(i)] = strings.Replace(tagParts[fmt.Sprint(i)], "\\\\", "\\", -1)
				if skipNext {
					i++
				}
			}
		}

		// Reconstruct list after rejoining
		for _, v := range tagParts {
			tagList = strings.SplitN(v, ":", 2)
			tagMap[tagList[0]] = ""
			if len(tagList) > 1 {
				tagMap[tagList[0]] = tagList[1]
			}
		}

		// Process uadmin tag
		// First, get the fields meta properties
		f.Help = tagMap["help"]
		f.Pattern = tagMap["pattern"]
		f.PatternMsg = tagMap["pattern_msg"]
		_, f.Required = tagMap["required"]
		_, f.Hidden = tagMap["hidden"]
		f.FormDisplay = !f.Hidden
		_, f.Encrypt = tagMap["encrypt"]
		_, f.Approval = tagMap["approval"]
		_, f.WebCam = tagMap["webcam"]
		_, f.Stringer = tagMap["stringer"]
		f.Min = tagMap["min"]
		f.Max = tagMap["max"]
		f.Format = tagMap["format"]
		f.DefaultValue = tagMap["default_value"]
		_, f.Searchable = tagMap["search"]
		//_, f.LimitChoicesTo = tagMap["limit_choices_to"]
		_, f.ListDisplay = tagMap["list_exclude"]
		f.ListDisplay = !f.ListDisplay
		if val, ok := tagMap["read_only"]; ok {
			if val == "" || val == cTRUE {
				f.ReadOnly = cTRUE
			} else {
				if !strings.HasPrefix(val, "true,") {
					val = "true," + val
				}
				f.ReadOnly = val
			}
		}
		_, f.Filter = tagMap["filter"]
		_, f.CategoricalFilter = tagMap["categorical_filter"]

		// Get custom display name of the field
		if val, ok := tagMap["display_name"]; ok {
			f.DisplayName = val
		}

		// Get the type name
		// f.TypeName = t.Field(index).Type.Name()
		typeName := strings.Split(t.Field(index).Type.String(), ".")
		f.TypeName = typeName[len(typeName)-1]

		// Process the field's data type
		if t.Field(index).Type == SType {
			f.Type = "string"
		}
		if t.Field(index).Type == DType || t.Field(index).Type == DType1 {
			f.Type = "date"
		}
		if t.Field(index).Type == BType {
			f.Type = "bool"
		}
		if t.Field(index).Type.Kind() == reflect.Struct && t.Field(index).Anonymous {
			f.Type = cID
			f.Name = "ID"
			f.DisplayName = "ID"
			f.ColumnName = "id"
		}
		if t.Field(index).Type == NType || t.Field(index).Type == NType1 || t.Field(index).Type == NType2 || t.Field(index).Type == NType3 || t.Field(index).Type == NType4 || t.Field(index).Type == NType5 || t.Field(index).Type == NType6 || t.Field(index).Type == NType7 {
			f.Type = "number"
		}
		if (t.Field(index).Type.Kind() == reflect.Struct && !t.Field(index).Anonymous && t.Field(index).Type != DType) ||
			(t.Field(index).Type.Kind() == reflect.Ptr && t.Field(index).Type.Elem().Kind() == reflect.Struct && !t.Field(index).Anonymous && t.Field(index).Type != DType1) {
			f.Type = cFK
			if val, ok := t.FieldByName(t.Field(index).Name + "ID"); ok {
				// Check if the FK field is a number
				if !(val.Type == NType || val.Type == NType1 || val.Type == NType2 || val.Type == NType3 || val.Type == NType4 || val.Type == NType5) {
					Trail(ERROR, "Invalid FK %s.%s your %sID field is not an integer based number", t.Name(), t.Field(index).Name, t.Field(index).Name)
				}
			} else {
				Trail(ERROR, "Invalid FK %s.%s no ID field found. Please add %sID field with a number type to your struct", t.Name(), t.Field(index).Name, t.Field(index).Name)
			}
		}
		if f.Type == cNUMBER && strings.HasSuffix(t.Field(index).Name, "ID") && t.Field(index).Name != "ID" {
			if _, ok := t.FieldByName(strings.TrimSuffix(t.Field(index).Name, "ID")); ok {
				continue
			}
		}
		if t.Field(index).Type.Kind() == reflect.Slice {
			f.Type = cM2M
		}

		// End of basic type, now we process extended types
		// First string extended type
		if _, ok := tagMap[cEMAIL]; ok {
			if f.Type != cSTRING {
				Trail(WARNING, "Invalid email tag in %s.%s, field data type shold be string not (%s)", s.Name, f.Name, f.Type)
			} else {
				f.Type = cEMAIL
			}
		}

		if _, ok := tagMap[cMULTILINGUAL]; ok {
			if f.Type != cSTRING {
				Trail(WARNING, "Invalid multilingual tag in %s.%s, field data type shold be string not (%s).", s.Name, f.Name, f.Type)
			} else {
				f.Type = cMULTILINGUAL

				for _, lang := range activeLangs {
					f.Translations = append(f.Translations, translation{
						Name:    fmt.Sprintf("%s (%s)", lang.Name, lang.EnglishName),
						Code:    lang.Code,
						Flag:    lang.Flag,
						Default: lang.Default,
						Active:  lang.Active,
					})
				}
			}
		}
		if _, ok := tagMap[cIMAGE]; ok {
			if f.Type != cSTRING {
				Trail(WARNING, "Invalid image tag in %s.%s, field data type shold be string not (%s).", s.Name, f.Name, f.Type)
			} else {
				f.Type = cIMAGE
				f.UploadTo = tagMap["upload_to"]
				if f.UploadTo != "" {
					if f.UploadTo[0] != '/' {
						f.UploadTo = "/" + f.UploadTo
					}
					if f.UploadTo[len(f.UploadTo)-1] != '/' {
						f.UploadTo = f.UploadTo + "/"
					}
				}
			}
		}
		if _, ok := tagMap[cFILE]; ok {
			if f.Type != cSTRING {
				Trail(WARNING, "Invalid file tag in %s.%s, field data type shold be string not (%s).", s.Name, f.Name, f.Type)
			} else {
				f.Type = cFILE
				f.UploadTo = tagMap["upload_to"]
				if f.UploadTo != "" {
					if f.UploadTo[0] != '/' {
						f.UploadTo = "/" + f.UploadTo
					}
					if f.UploadTo[len(f.UploadTo)-1] != '/' {
						f.UploadTo = f.UploadTo + "/"
					}
				}
			}
		}
		if _, ok := tagMap[cPASSWORD]; ok {
			if f.Type != cSTRING {
				Trail(WARNING, "Invalid password tag in %s.%s, field data type shold be string not (%s).", s.Name, f.Name, f.Type)
			} else {
				f.Type = cPASSWORD
			}
		}
		if _, ok := tagMap[cHTML]; ok {
			if f.Type != cSTRING {
				Trail(WARNING, "Invalid html tag in %s.%s, field data type shold be string not (%s).", s.Name, f.Name, f.Type)
			} else {
				f.Type = cHTML
			}
		}
		if _, ok := tagMap[cLINK]; ok {
			if f.Type != cSTRING {
				Trail(WARNING, "Invalid link tag in %s.%s, field data type shold be string not (%s).", s.Name, f.Name, f.Type)
			} else {
				f.Type = cLINK
			}
		}
		if _, ok := tagMap[cCODE]; ok {
			if f.Type != cSTRING {
				Trail(WARNING, "Invalid code tag in %s.%s, field data type shold be string not (%s).", s.Name, f.Name, f.Type)
			} else {
				f.Type = cCODE
			}
		}

		// Now we process number extended types
		if val, ok := tagMap[cPROGRESSBAR]; ok {
			if f.Type != cNUMBER {
				Trail(WARNING, "Invalid progress_bar tag in %s.%s, field data type shold be number not (%s).", s.Name, f.Name, f.Type)
			} else if val == "" {
				// This is the case were the progress_bar tag was passed with no parameters
				// In this case we create a default progress bar from 0 to 100 and make the color blue
				// Tag Foramt: progress_bar
				f.Type = cPROGRESSBAR
				f.ProgressBar = map[float64]string{
					100: defaultProgressBarColor,
				}
			} else {
				// This is the base where a progress bar was passed with parameters
				// Allowed formats are:
				// progress_bar:100.0                      (Set max value)
				// progress_bar:100.0:#0f0                 (set max value and colors)
				// progress_bar:0.4:#f00,0.7:#ff0,1.0:#0f0 (set multiple colors and their thresholds)
				progressList := strings.Split(val, ",")
				if val, err := strconv.ParseFloat(progressList[0], 10); len(progressList) == 1 && err == nil {
					//TODO: Make default color adjustable system wide
					f.ProgressBar = map[float64]string{
						val: defaultProgressBarColor,
					}
					f.Type = cPROGRESSBAR
				} else {
					if len(progressList) == 1 && !strings.Contains(progressList[0], ":") {
						Trail(WARNING, "Invalid progress_bar tag in %s.%s, unknown single value format (%s)", s.Name, f.Name, progressList[0])
					} else {
						errorFound := false
						f.ProgressBar = map[float64]string{}
						for _, v := range progressList {
							thresholdList := strings.Split(v, ":")
							// TODO: Trim white space for thresholdList
							if len(thresholdList) != 2 {
								Trail(WARNING, "Invalid progress_bar tag in %s.%s, unknown multi value format (%s)", s.Name, f.Name, v)
								errorFound = true
							} else if _, err := strconv.ParseFloat(thresholdList[0], 10); err != nil {
								Trail(WARNING, "Invalid progress_bar tag in %s.%s, invalid number for threshold (%s)", s.Name, f.Name, v)
								errorFound = true
								//TODO: Check the color
							} else {
								val, _ := strconv.ParseFloat(thresholdList[0], 10)
								f.ProgressBar[val] = thresholdList[1]
							}
						}
						if !errorFound {
							f.Type = cPROGRESSBAR
						}
					}
				}
			}
		}

		if _, ok := tagMap["money"]; ok {
			if f.Type != cNUMBER {
				Trail(WARNING, "Invalid money tag in %s.%s, field data type shold be number not (%s).", s.Name, f.Name, f.Type)
			} else {
				f.Type = cMONEY
			}
		}

		// Process static list type
		// The way this is checked is if the type is not an int and the kind is an int the
		// it is a static list
		if t.Field(index).Type != NType && t.Field(index).Type.Kind() == reflect.Int {
			f.Type = cLIST

			f.Choices = []Choice{
				{" - ", 0, false},
			}
			for i := 0; i < t.Field(index).Type.NumMethod(); i++ {
				v := t.Field(index).Type.Method(i).Name
				e := reflect.ValueOf(a).Field(index)
				e1 := reflect.Indirect(e).Method(i)

				tempK := e1.Call([]reflect.Value{})
				k, err := strconv.ParseUint(fmt.Sprint(tempK[0]), 10, 64)
				if err != nil {
					Trail(ERROR, "Unable to get list value for %s.%s because %s", modelName, f.Name, err.Error())
				}

				// TODO: Make list multi lingual
				f.Choices = append(f.Choices, Choice{
					V: getDisplayName(v),
					K: uint(k),
				})
			}
		}
		f.FormDisplay = true
		s.Fields = append(s.Fields, f)
	}

	// Add method Fields
	for index := 0; index < t.NumMethod(); index++ {
		// Check if the method should be included in the field list
		if strings.Contains(t.Method(index).Name, "__Form") || strings.Contains(t.Method(index).Name, "__List") {

			if strings.ToLower(string(t.Method(index).Name[0])) == string(t.Method(index).Name[0]) {
				continue
			}
			f := F{
				Translations: []translation{},
			}
			f.Name = t.Method(index).Name
			f.DisplayName = strings.TrimSuffix(t.Method(index).Name, "__Form")
			f.DisplayName = strings.TrimSuffix(f.DisplayName, "__List")
			f.DisplayName = strings.TrimSuffix(f.DisplayName, "__Form")
			f.DisplayName = getDisplayName(f.DisplayName)
			f.Type = cSTRING
			f.ReadOnly = cTRUE
			f.IsMethod = true
			f.FormDisplay = strings.Contains(t.Method(index).Name, "__Form")
			f.ListDisplay = strings.Contains(t.Method(index).Name, "__List")
			s.Fields = append(s.Fields, f)
		}
	}

	// Initialize lists
	s.IncludeFormJS = []string{}
	s.IncludeListJS = []string{}

	//Schema[strings.ToLower(t.Name())] = s
	return s, true
}
