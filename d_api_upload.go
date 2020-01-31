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
		// Process File
		// Check if the file is type file or image
		var field *F
		for i := range schema.Fields {
			if schema.Fields[i].ColumnName == k[1:] {
				field = &schema.Fields[i]
				r.MultipartForm.File[k[1:]] = r.MultipartForm.File[k]
				break
			}
		}
		if field == nil {
			Trail(WARNING, "dAPIUpload received a file that has no field: %s", k)
			continue
		}

		s := r.Context().Value(CKey("session"))
		var session *Session
		if s != nil {
			session = s.(*Session)
		}

		fileName := processUpload(r, field, schema.ModelName, session, schema)
		if fileName != "" {
			fileList[field.ColumnName] = fileName
		}
	}
	return fileList, nil
}
