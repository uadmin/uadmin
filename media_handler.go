package uadmin

import (
	"net/http"
	"os"
	"path"
	"strings"
)

func MediaHandler(w http.ResponseWriter, r *http.Request, relPath string) {
	// session := IsAuthenticated(r)
	// if session == nil && !PublicMedia {
	// 	loginHandler(w, r)
	// 	return
	// }

	// r.URL.Path = strings.TrimPrefix(r.URL.Path, "/media/")
	// file, err := os.Open("./media/" + path.Clean(r.URL.Path))
	// if err != nil {
	// 	pageErrorHandler(w, r, session)
	// 	return
	// }
	// io.Copy(w, file)
	// file.Close()

	pathCorrected := strings.TrimPrefix(r.URL.Path, relPath)
	fName := path.Clean(pathCorrected)

	f, err := os.Open("." + fName)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	defer f.Close()
	stat, err := os.Stat("." + fName)
	if err != nil || stat.IsDir() {
		w.WriteHeader(404)
		return
	}
	modTime := stat.ModTime()
	if RetainMediaVersions {
		w.Header().Add("Cache-Control", "private, max-age=604800")
	} else {
		w.Header().Add("Cache-Control", "private, max-age=3600")
	}

	http.ServeContent(w, r, "."+fName, modTime, f)

	// Delete the file if exported to excel
	if strings.HasPrefix(fName, "/media/export/") {
		filePart := strings.TrimPrefix(fName, "/media/export/")
		filePart = path.Clean(filePart)
		if filePart != "" && !strings.HasSuffix(filePart, "index.html") {
			os.Remove("./media/export/" + filePart)
		}
	}
}

func mediaHandler(w http.ResponseWriter, r *http.Request) {
	session := IsAuthenticated(r)
	if session == nil && !PublicMedia {
		loginHandler(w, r)
		return
	}

	// r.URL.Path = strings.TrimPrefix(r.URL.Path, "/media/")
	// file, err := os.Open("./media/" + path.Clean(r.URL.Path))
	// if err != nil {
	// 	pageErrorHandler(w, r, session)
	// 	return
	// }
	// io.Copy(w, file)
	// file.Close()

	fName := path.Clean(r.URL.Path)

	f, err := os.Open("." + fName)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	defer f.Close()
	stat, err := os.Stat("." + fName)
	if err != nil || stat.IsDir() {
		w.WriteHeader(404)
		return
	}
	modTime := stat.ModTime()
	if RetainMediaVersions {
		w.Header().Add("Cache-Control", "private, max-age=604800")
	} else {
		w.Header().Add("Cache-Control", "private, max-age=3600")
	}

	http.ServeContent(w, r, "."+fName, modTime, f)

	// Delete the file if exported to excel
	if strings.HasPrefix(fName, "/media/export/") {
		filePart := strings.TrimPrefix(fName, "/media/export/")
		filePart = path.Clean(filePart)
		if filePart != "" && !strings.HasSuffix(filePart, "index.html") {
			os.Remove("./media/export/" + filePart)
		}
	}
}
