package uadmin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

/*
// Field appears in JSON as key "myName".
Field int `json:"myName"`

// Field appears in JSON as key "myName" and
// the field is omitted from the object if its value is empty,
// as defined above.
Field int `json:"myName,omitempty"`

// Field appears in JSON as key "Field" (the default), but
// the field is skipped if empty.
// Note the leading comma.
Field int `json:",omitempty"`

// Field is ignored by this package.
Field int `json:"-"`

// Field appears in JSON as key "-".
Field int `json:"-,"`
*/

type fieldLanguage struct {
	DisplayName       string `json:"display_name"`
	updateDisplayName bool
	Help              string `json:"help"`
	updateHelp        bool
	PatternMsg        string `json:"pattern_msg"`
	updatePatternMsg  bool
	ErrMsg            map[string]string `json:"err_msg"`
	Choices           map[int]string
	updateChoice      []bool
}

type structLanguage struct {
	DisplayName       string `json:"display_name"`
	updateDisplayName bool
	Fields            map[string]fieldLanguage `json:"fields"`
}

const translateMe = "Translate me ---> "

// syncCustomTranslation is a function for creating and updating custom translation files
func syncCustomTranslation(path string) map[string]int {
	var err error
	var buf []byte
	stat := map[string]int{}

	pathParts := strings.Split(path, "/")
	if len(pathParts) != 2 {
		Trail(ERROR, "Custom translation file path is incorrect (%s)", path)
		return stat
	}
	group := pathParts[0]
	name := pathParts[1]

	os.MkdirAll("./static/i18n/"+group+"/", 0744)
	fileName := "./static/i18n/" + group + "/" + name + ".en.json"
	langMap := map[string]string{}
	if _, err = os.Stat(fileName); os.IsNotExist(err) {
		ioutil.WriteFile(fileName, []byte{'{', '}'}, 0644)
		stat["en"] = 0
	} else {
		buf, err = ioutil.ReadFile(fileName)
		if err != nil {
			Trail(ERROR, "Unable to read system translation file (%s)", fileName)
			return stat
		}
		err = json.Unmarshal(buf, &langMap)
		if err != nil {
			Trail(ERROR, "Invalid format of system translation file (%s). %s", fileName, err)
		}
		for _, lang := range activeLangs {
			if lang.Code == "en" {
				stat["en"] = len(langMap)
				continue
			}
			stat[lang.Code] = 0
			updateRequired := false
			langFileName := "./static/i18n/" + group + "/" + name + "." + lang.Code + ".json"
			langSystemMap := map[string]string{}
			if _, err = os.Stat(langFileName); os.IsNotExist(err) {
				ioutil.WriteFile(langFileName, []byte{'{', '}'}, 0644)
			}
			buf, err = ioutil.ReadFile(langFileName)
			if err != nil {
				Trail(ERROR, "Unable to read system translation file (%s)", langFileName)
				return stat
			}
			err = json.Unmarshal(buf, &langSystemMap)
			if err != nil {
				Trail(ERROR, "Invalid format of system translation file (%s). %s", langFileName, err)
			}
			for k, v := range langMap {
				if val, ok := langSystemMap[k]; !ok {
					updateRequired = true
					langSystemMap[k] = translateMe + v
				} else {
					if !strings.HasPrefix(val, translateMe) {
						stat[lang.Code]++
					}
				}
			}
			if updateRequired {
				saveLangFile(langSystemMap, langFileName)
			}
		}
	}
	return stat
}

func syncModelTranslation(m ModelSchema) map[string]int {
	var err error
	var buf []byte
	stat := map[string]int{
		"en": 1,
	}

	// Generate/Sync original language file
	// First parse the schema into a structLanguage
	structLang := structLanguage{}
	structLang.DisplayName = m.DisplayName
	structLang.Fields = map[string]fieldLanguage{}
	modelCount := 1
	fileCount := 1
	for _, f := range m.Fields {
		structLang.Fields[f.Name] = fieldLanguage{
			DisplayName:  f.DisplayName,
			Help:         f.Help,
			PatternMsg:   f.PatternMsg,
			ErrMsg:       map[string]string{},
			Choices:      map[int]string{},
			updateChoice: []bool{},
		}
		if f.DisplayName != "" {
			modelCount++
		}
		if f.Help != "" {
			modelCount++
		}
		if f.Help != "" {
			modelCount++
		}
	}

	// Get information about the model
	pkgName := fmt.Sprint(reflect.TypeOf(models[m.ModelName]))
	pkgName = strings.Split(pkgName, ".")[0]

	// Get the model's original language file
	os.MkdirAll("./static/i18n/"+pkgName+"/", 0744)
	fileName := "./static/i18n/" + pkgName + "/" + m.ModelName + ".en.json"

	// Check if the fist doesn't exist and create it
	if _, err = os.Stat(fileName); os.IsNotExist(err) {
		buf, _ = json.MarshalIndent(structLang, "", "  ")
		err = ioutil.WriteFile(fileName, buf, 0644)
		if err != nil {
			Trail(ERROR, "generateTranslation error writing a file %v", err)
		}
		fileCount = modelCount
	} else {
		// It is exists, read it
		buf, err = ioutil.ReadFile(fileName)
		if err != nil {
			Trail(ERROR, "Unable to read language file (%s)", fileName)
		}

		// Parse the on disk file to a structLanguage
		langOnFile := structLanguage{}
		err = json.Unmarshal(buf, &langOnFile)
		if err != nil {
			Trail(ERROR, "Invalid format for language file (%s)", fileName)
		}

		// Check if there are any changes that required updating the language file
		requiresUpdate := false
		if langOnFile.DisplayName != structLang.DisplayName {
			requiresUpdate = true
			langOnFile.updateDisplayName = true
			langOnFile.DisplayName = structLang.DisplayName
		}
		if langOnFile.DisplayName != "" {
			fileCount++
		}

		for k, v := range structLang.Fields {
			// Check if the field is not on file
			if onFileV, ok := langOnFile.Fields[k]; !ok {
				requiresUpdate = true
				langOnFile.Fields[k] = v
				fileCount++
			} else {
				// If the field is on disk, then verify all variables
				if v.DisplayName != onFileV.DisplayName {
					requiresUpdate = true
					onFileV.updateDisplayName = true
					onFileV.DisplayName = v.DisplayName
				}
				if v.PatternMsg != onFileV.PatternMsg {
					requiresUpdate = true
					onFileV.updatePatternMsg = true
					onFileV.PatternMsg = v.PatternMsg
				}
				if v.Help != onFileV.Help {
					requiresUpdate = true
					onFileV.updateHelp = true
					onFileV.Help = v.Help
				}
				if onFileV.DisplayName != "" {
					fileCount++
				}
				if onFileV.Help != "" {
					fileCount++
				}
				if onFileV.PatternMsg != "" {
					fileCount++
				}
				fileCount += len(onFileV.ErrMsg)

				// Assign back the onFileV
				langOnFile.Fields[k] = onFileV
				// Assign back the fieldLang
				structLang.Fields[k] = onFileV
			}
		}

		// If the file was changed, write it back to disk
		if requiresUpdate {
			buf, _ = json.MarshalIndent(langOnFile, "", "  ")
			err = ioutil.WriteFile(fileName, buf, 0644)
			if err != nil {
				Trail(ERROR, "Unable to write language file (%s)", fileName)
				return stat
			}
		}

		// Finally update the model's structLanguage with the updated one from disk
		structLang = langOnFile
	}
	stat["en"] = fileCount

	// Sync active languages
	for _, lang := range activeLangs {
		if lang.Code == "en" {
			continue
		}
		updateRequired := false
		stat[lang.Code] = 0

		// The active language file name
		langFileName := "./static/i18n/" + pkgName + "/" + m.ModelName + "." + lang.Code + ".json"
		structLangOnFile := structLanguage{}

		// Check if the language file exists
		if _, err = os.Stat(langFileName); os.IsNotExist(err) {
			buf, _ = json.MarshalIndent(structLangOnFile, "", "  ")
			ioutil.WriteFile(langFileName, buf, 0644)
		}

		// Read/Parse language file from disk
		buf, err = ioutil.ReadFile(langFileName)
		if err != nil {
			Trail(ERROR, "Unable to read system translation file (%s)", langFileName)
			continue
		}
		err = json.Unmarshal(buf, &structLangOnFile)
		if err != nil {
			Trail(ERROR, "Invalid format of system translation file (%s). %s", langFileName, err)
			continue
		}

		// If language file is empty then initialize it
		if structLangOnFile.DisplayName == "" {
			updateRequired = true
			structLangOnFile.DisplayName = translateMe + m.DisplayName
			structLangOnFile.Fields = map[string]fieldLanguage{}
		} else {
			// if the file exists then verify its content
			if structLang.updateDisplayName {
				updateRequired = true
				structLangOnFile.DisplayName = translateMe + m.DisplayName
			}
			if structLangOnFile.DisplayName != "" && !strings.HasPrefix(structLangOnFile.DisplayName, translateMe) {
				stat[lang.Code]++
			}
		}

		for k, v := range structLang.Fields {
			// Check if the field is not on file and add it
			if langV, ok := structLangOnFile.Fields[k]; !ok {
				updateRequired = true
				langV = fieldLanguage{
					ErrMsg: map[string]string{},
				}
				langV.DisplayName = translateMe + v.DisplayName
				if v.Help != "" {
					langV.Help = translateMe + v.Help
				}
				if v.PatternMsg != "" {
					langV.PatternMsg = translateMe + v.PatternMsg
				}
				for errK, errV := range v.ErrMsg {
					langV.ErrMsg[errK] = translateMe + errV
				}
				structLangOnFile.Fields[k] = langV
			} else {
				// If the field exists, then verify it's contents
				if v.updateDisplayName {
					updateRequired = true
					langV.DisplayName = translateMe + v.DisplayName
				}
				if langV.DisplayName != "" && !strings.HasPrefix(langV.DisplayName, translateMe) {
					stat[lang.Code]++
				}
				if v.updateHelp {
					updateRequired = true
					langV.Help = translateMe + v.Help
				}
				if langV.Help != "" && !strings.HasPrefix(langV.Help, translateMe) {
					stat[lang.Code]++
				}
				if v.updatePatternMsg {
					updateRequired = true
					langV.PatternMsg = translateMe + v.PatternMsg
				}
				if langV.PatternMsg != "" && !strings.HasPrefix(langV.PatternMsg, translateMe) {
					stat[lang.Code]++
				}
				for errK, errV := range v.ErrMsg {
					if _, ok = langV.ErrMsg[errK]; !ok {
						updateRequired = true
						langV.ErrMsg[errK] = translateMe + errV
					}
					if langV.ErrMsg[errK] != "" && !strings.HasPrefix(langV.ErrMsg[errK], translateMe) {
						stat[lang.Code]++
					}
				}

				// Assign the field back into the structLanguage
				structLangOnFile.Fields[k] = langV
			}
		}

		// If there are any changes, write the file back to disk
		if updateRequired {
			saveLangFile(structLangOnFile, langFileName)
			if err != nil {
				Trail(ERROR, "Unable to write language file language file (%s)", langFileName)
			}
		}
	}
	return stat
}

func saveLangFile(v interface{}, fileName string) {
	buf, _ := json.MarshalIndent(v, "", "  ")
	buf = bytes.Replace(buf, []byte("\\u003c"), []byte("<"), -1)
	buf = bytes.Replace(buf, []byte("\\u003e"), []byte(">"), -1)
	langMapCache[fileName] = buf
	ioutil.WriteFile(fileName, buf, 0644)
}
