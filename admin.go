package uadmin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

/*
	admin Tags
	read_only:TRUE
	email:TRUE
	hidden:TRUE
	html:TRUE
	fk:"ModelName"
	list:TRUE
	list_filter:TRUE
	search:TRUE
	dontCache:TRUE
  	required:TRUE
  	help:TRUE
  	pattern:TRUE
  	pattern_msg:"Message"
	max:"int"
	min:"int"
	link:TRUE
	file:TRUE
	dependsOn:""
	linkerObj:""
	linkerParentField:""
	linkerChildField:""
	childObj:""
	upload_to:"path"
	code:"true"
	money:"true" use on float
	defaultValue:""
*/

// commaf is a function to format number with thousand separator
// and two decimal points
func commaf(j interface{}) string {
	v, _ := strconv.ParseFloat(fmt.Sprint(j), 64)
	buf := &bytes.Buffer{}
	if v < 0 {
		buf.Write([]byte{'-'})
		v = 0 - v
	}
	s := fmt.Sprintf("%.2f", v)

	comma := []byte{','}

	parts := strings.Split(s, ".")
	pos := 0
	if len(parts[0])%3 != 0 {
		pos += len(parts[0]) % 3
		buf.WriteString(parts[0][:pos])
		buf.Write(comma)
	}
	for ; pos < len(parts[0]); pos += 3 {
		buf.WriteString(parts[0][pos : pos+3])
		buf.Write(comma)
	}
	buf.Truncate(buf.Len() - 1)

	if len(parts) > 1 {
		buf.Write([]byte{'.'})
		buf.WriteString(parts[1])
	}
	return buf.String()
}

func isLocal(Addr string) bool {
	if strings.Contains(Addr, ":") && !strings.Contains(Addr, ".") {
		Addr = strings.TrimPrefix(Addr, "[")
		if strings.HasPrefix(Addr, "::") || strings.HasPrefix(Addr, "fc") || strings.HasPrefix(Addr, "fd") {
			return true
		}
	}
	p := strings.Split(strings.Split(Addr, ":")[0], ".")
	if len(p) != 4 {
		return false
	}
	_, err := strconv.ParseInt(p[2], 10, 64)
	if err != nil {
		return false
	}
	_, err = strconv.ParseInt(p[3], 10, 64)
	if err != nil {
		return false
	}
	v1, err := strconv.ParseInt(p[0], 10, 64)
	if err != nil {
		return false
	}
	v2, err := strconv.ParseInt(p[1], 10, 64)
	if err != nil {
		return false
	}
	if v1 == 10 {
		return true
	}
	if v1 == 172 {
		if v2 >= 16 && v2 <= 31 {
			return true
		}
	}
	if v1 == 192 && v2 == 168 {
		return true
	}
	if v1 == 127 {
		return true
	}
	return false
}

// saver is an interface to deal with form froms
type saver interface {
	Save()
}

// counter !
type counter interface {
	Count(interface{}, interface{}, ...interface{}) int
}

type adminPager interface {
	AdminPage(string, bool, int, int, interface{}, interface{}, ...interface{}) error
}

func paginationHandler(itemCount int, PageLength int) (i int) {
	pCount := (float64(itemCount) / float64(PageLength))
	if pCount > float64(int(pCount)) {
		pCount++
	}
	i = int(pCount)
	if i == 1 {
		i--
	}
	return i
}

// toSnakeCase !
func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// JSONMarshal Generates JSON format from an object
func JSONMarshal(v interface{}, safeEncoding bool) ([]byte, error) {
	// b, err := json.Marshal(v)
	b, err := jsonMarshal(v)

	if safeEncoding {
		b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
		b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
		b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
	}
	return b, err
}

func nullZeroValueStructs(record map[string]interface{}) map[string]interface{} {
	for k := range record {
		switch v := record[k].(type) {
		case map[string]interface{}:
			if id, ok := v["ID"].(float64); ok && id == 0 {
				record[k] = nil
			} else if id, ok := v["id"].(float64); ok && id == 0 {
				record[k] = nil
			} else {
				record[k] = nullZeroValueStructs(v)
			}
		}
	}
	return record
}

func removeZeroValueStructs(buf []byte) []byte {
	response := map[string]interface{}{}
	json.Unmarshal(buf, &response)
	if val, ok := response["result"].(map[string]interface{}); ok {
		val = nullZeroValueStructs(val)
		buf, _ = json.Marshal(val)
		return buf
	}
	if _, ok := response["result"].([]interface{}); !ok {
		return buf
	}
	val := response["result"].([]interface{})
	var record map[string]interface{}
	for i := range val {
		record = val[i].(map[string]interface{})
		record = nullZeroValueStructs(record)
		val[i] = record
	}
	response["result"] = val
	buf, _ = json.Marshal(response)
	return buf
}

func jsonMarshal(v interface{}) ([]byte, error) {
	var buf []byte
	var err error
	if CompressJSON {
		buf, err = json.Marshal(v)
		if err == nil && RemoveZeroValueJSON {
			buf = removeZeroValueStructs(buf)
		}
	} else {
		buf, err = json.MarshalIndent(v, "", " ")
	}

	return buf, err
}

// ReturnJSON returns json to the client
func ReturnJSON(w http.ResponseWriter, r *http.Request, v interface{}) {
	// Set content type in header
	w.Header().Set("Content-Type", "application/json")

	// Marshal content
	b, err := jsonMarshal(v)
	if err != nil {
		response := map[string]interface{}{
			"status":    "error",
			"error_msg": fmt.Sprintf("unable to encode JSON. %s", err),
		}
		b, _ = jsonMarshal(response)
		w.Write(b)
		return
	}

	if CustomizeJSON != nil {
		b = CustomizeJSON(w, r, v, b)
	}

	w.Write(b)
}

func stripHTMLScriptTag(v string) string {
	doc, err := html.Parse(strings.NewReader(v))
	if err != nil {
		return ""
	}
	removeScript(doc)
	b := bytes.NewBuffer([]byte{})
	if err := html.Render(b, doc); err != nil {
		return ""
	}
	return b.String()
}

func removeScript(n *html.Node) {
	// if note is script tag
	if n.Type == html.ElementNode && (strings.Contains(n.Data, "script") || strings.Contains(n.Data, "frame")) {
		n.Parent.RemoveChild(n)
		return // script tag is gone...
	}
	// traverse DOM
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		removeScript(c)
	}
}
