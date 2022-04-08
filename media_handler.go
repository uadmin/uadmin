package uadmin

import (
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

func mediaHandler(w http.ResponseWriter, r *http.Request) {
	session := IsAuthenticated(r)
	if session == nil && !PublicMedia {
		loginHandler(w, r)
		return
	}

	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/media/")
	file, err := os.Open("./media/" + path.Clean(r.URL.Path))
	if err != nil {
		pageErrorHandler(w, r, session)
		return
	}
	io.Copy(w, file)
	file.Close()

	// Delete the file if exported to excel
	if strings.HasPrefix(r.URL.Path, "export/") {
		filePart := strings.TrimPrefix(r.URL.Path, "export/")
		filePart = path.Clean(filePart)
		if filePart != "" && !strings.HasSuffix(filePart, "index.html") {
			os.Remove("./media/export/" + filePart)
		}
	}
}
