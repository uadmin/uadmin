package uadmin

import (
	"fmt"
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

	title, _, _ := getHTMLTag(w.Result().Body, "title")
	if title != "uAdmin - Login" {
		t.Errorf("Invalid page returned. Expected Login, got (%s)", title)
	}
}

func traverse(n *html.Node, tag string) (string, map[string]string, bool) {
	if isTagElement(n, tag) {
		tempMap := map[string]string{}
		for i := range n.Attr {
			tempMap[n.Attr[i].Key] = n.Attr[i].Val
		}
		if n.FirstChild == nil {
			return "", tempMap, true
		}
		return strings.TrimSpace(n.FirstChild.Data), tempMap, true
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result, attr, ok := traverse(c, tag)
		if ok {
			return result, attr, ok
		}
	}

	return "", map[string]string{}, false
}

func getHTMLTag(r io.Reader, tag string) (string, map[string]string, bool) {
	doc, err := html.Parse(r)
	if err != nil {
		Trail(ERROR, "Fail to parse html")
		return "", map[string]string{}, false
	}

	return traverse(doc, tag)
}

func isTagElement(n *html.Node, tag string) bool {
	return n.Type == html.ElementNode && n.Data == tag
}

func tagSearch(n *html.Node, tag string, path string, index int) ([]string, []string, []map[string]string) {
	paths := []string{}
	content := []string{}
	attr := []map[string]string{}

	if path == "" {
		if n.Data != "" {
			path = fmt.Sprintf("%s[%d]", n.Data, index)
		}
	} else {
		path = path + "/" + fmt.Sprintf("%s[%d]", n.Data, index)
	}

	if isTagElement(n, tag) {
		if n.FirstChild == nil {
			content = append(content, "")
		} else {
			content = append(content, strings.TrimSpace(n.FirstChild.Data))
		}
		paths = append(paths, path)
		tempMap := map[string]string{}
		for i := range n.Attr {
			tempMap[n.Attr[i].Key] = n.Attr[i].Val
		}
		attr = append(attr, tempMap)
	}

	index = 0
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		childPaths, childContent, childAttr := tagSearch(c, tag, path, index)
		paths = append(paths, childPaths...)
		content = append(content, childContent...)
		attr = append(attr, childAttr...)
		if c.Type == html.ElementNode {
			index++
		}
	}
	return paths, content, attr
}

func getHTMLTagList(r io.Reader, tag string) (paths []string, content []string, attr []map[string]string) {
	doc, err := html.Parse(r)
	if err != nil {
		Trail(ERROR, "Failed to parse html")
		return
	}
	return tagSearch(doc, tag, "", 0)
}

func parseHTML(r io.Reader, t *testing.T) (*html.Node, error) {
	doc, err := html.Parse(r)
	if err != nil {
		t.Errorf("Unable to parse html stream")
	}
	return doc, err
}
