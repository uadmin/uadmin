package uadmin

import (
	"net/http"
)

func dAPIUpload(w http.ResponseWriter, r *http.Request, schema *ModelSchema) (map[string]string, error) {
	fileList := map[string]string{}

	if r.MultipartForm == nil {
		return fileList, nil
	}

	// make a list of files
	kList := []string{}
	for k := range r.MultipartForm.File {
		kList = append(kList, k)
	}

	for _, k := range kList {
		// Process File
		var field *F = schema.FieldByColumnName(k[1:])
		if field == nil {
			Trail(WARNING, "dAPIUpload received a file that has no field: %s", k)
			continue
		}

		r.MultipartForm.File[k[1:]] = r.MultipartForm.File[k]

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
