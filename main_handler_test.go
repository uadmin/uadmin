package uadmin

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

// TestMainHandler is a unit testing function for mainHandler() function
func TestMainHandler(t *testing.T) {
	r := httptest.NewRequest("GET", "http://0.0.0.0:5000/", nil)
	w := httptest.NewRecorder()

	mainHandler(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Invalid code on requesting /. %d", w.Code)
	}

	title, _ := getHTMLTitle(w.Result().Body)
	title = strings.TrimSpace(title)
	if title != "uAdmin - Login" {
		t.Errorf("Invalid page returned. Expected Login, got (%s)", title)
	}
}

func isTitleElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}

func traverse(n *html.Node) (string, bool) {
	if isTitleElement(n) {
		return n.FirstChild.Data, true
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result, ok := traverse(c)
		if ok {
			return result, ok
		}
	}

	return "", false
}

func getHTMLTitle(r io.Reader) (string, bool) {
	doc, err := html.Parse(r)
	if err != nil {
		panic("Fail to parse html")
	}

	return traverse(doc)
}
