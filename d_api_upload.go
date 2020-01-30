package uadmin

import (
	"net/http"
)

func dAPIUpload(w http.ResponseWriter, r *http.Request, schema *ModelSchema) (map[string]string, error) {
	fileList := map[string]string{}

	if r.MultipartForm == nil {
		return fileList, nil
	}

	for k := range r.MultipartForm.File {
		Trail(DEBUG, "file: %s", k)
		// Process File
		// Check if the file is type file or image
		var field *F
		for i := range schema.Fields {
			if schema.Fields[i].ColumnName == k {
				field = &schema.Fields[i]
				break
			}
		}
		if field == nil {
			Trail(DEBUG, "no field: %s", k)
			continue
		}

		s := r.Context().Value(CKey("session"))
		var session *Session
		if s != nil {
			session = s.(*Session)
		}

		fileName := processUpload(r, field, schema.ModelName, session, schema)
		Trail(DEBUG, "fileName: %s", fileName)
		if fileName != "" {
			fileList[field.ColumnName] = fileName
		}
	}
	return fileList, nil
}
