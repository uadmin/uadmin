package uadmin

import (
	"io/ioutil"
	"math"
	"net/http/httptest"
	"testing"
)

// TestIsLocal is a unit testing function for isLocal() function
func TestIsLocal(t *testing.T) {
	examples := []struct {
		ip    string
		local bool
	}{
		{"127.0.0.1", true},
		{"127.1.1.1", true},
		{"128.0.0.1", false},
		{"192.168.0.1", true},
		{"192.168.1.1", true},
		{"192.168.1.255", true},
		{"192.168.100.1", true},
		{"192.168.200.1", true},
		{"10.0.0.1", true},
		{"10.0.1.1", true},
		{"10.1.1.1", true},
		{"10.100.0.1", true},
		{"10.100.100.1", true},
		{"10.100.100.100", true},
		{"11.0.0.1", false},
		{"8.8.8.8", false},
		{"8.8.4.4", false},
		{"1.1.1.1", false},
		{"1.1.1.100", false},
		{"172.16.0.1", true},
		{"172.17.0.1", true},
		{"172.18.0.1", true},
		{"172.19.0.1", true},
		{"172.20.0.1", true},
		{"172.21.0.1", true},
		{"172.22.0.1", true},
		{"172.23.0.1", true},
		{"172.24.0.1", true},
		{"172.25.0.1", true},
		{"172.26.0.1", true},
		{"172.27.0.1", true},
		{"172.28.0.1", true},
		{"172.29.0.1", true},
		{"172.30.0.1", true},
		{"172.31.0.1", true},
		{"172.15.0.1", false},
		{"172.32.0.1", false},
		{"[::1]", true},
		{"[::2]", true},
		{"[::f]", true},
		{"[::ffff]", true},
		{"[fc::1]", true},
		{"[fd::1]", true},
		{"[2400::1]", false},
		{"[2401::1]", false},
		{"[2401::100]", false},
		{"[2401::ffff]", false},
		{"[2401:1::ffff]", false},
		{"[2401:ffff:1::ffff]", false},
		{"a.32.0.1", false},
		{"172.a.0.1", false},
		{"172.32.a.1", false},
		{"172.32.0.a", false},
	}
	passedTests := 0
	for _, e := range examples {
		if isLocal(e.ip) != e.local {
			t.Errorf("isLocal(%s) = %v != %v", e.ip, isLocal(e.ip), e.local)
		} else {
			passedTests++
		}
	}
	Trail(OK, "Passed %d tests in TestIsLocal", passedTests)
	if passedTests < len(examples) {
		Trail(WARNING, "Failed %d tests in TestIsLocal", len(examples)-passedTests)
	}
}

// TestCommaf is a unit testing function for commaf() function
func TestCommaf(t *testing.T) {
	examples := []struct {
		in  interface{}
		out string
	}{
		{1, "1.00"},
		{-1, "-1.00"},
		{1.0, "1.00"},
		{-1.0, "-1.00"},
		{10, "10.00"},
		{100, "100.00"},
		{1000, "1,000.00"},
		{10.0, "10.00"},
		{100.0, "100.00"},
		{1000.0, "1,000.00"},
		{-10, "-10.00"},
		{-100, "-100.00"},
		{-1000, "-1,000.00"},
		{-10.0, "-10.00"},
		{-100.0, "-100.00"},
		{-1000.0, "-1,000.00"},
	}
	for _, e := range examples {
		if commaf(e.in) != e.out {
			t.Errorf("commaf(%s) = %s != %s", e.in, commaf(e.in), e.out)
		}
	}
}

// TestPaginationHandler is a unit testing function for paginationHandler() function
func TestPaginationHandler(t *testing.T) {
	examples := []struct {
		itemCount  int
		pageLength int
		pageCount  int
	}{
		{0, 100, 0},
		{1, 100, 0},
		{99, 100, 0},
		{100, 100, 0},
		{101, 100, 2},
		{1000, 100, 10},
		{1001, 100, 11},
		{0, 10, 0},
		{1, 10, 0},
		{8, 10, 0},
		{9, 10, 0},
		{10, 10, 0},
		{11, 10, 2},
		{19, 10, 2},
		{20, 10, 2},
		{21, 10, 3},
	}
	for _, e := range examples {
		if paginationHandler(e.itemCount, e.pageLength) != e.pageCount {
			t.Errorf("paginationHandler(%d, %d) = %d != %d", e.itemCount, e.pageLength, paginationHandler(e.itemCount, e.pageLength), e.pageCount)
		}
	}
}

// TestToSnakeCase is a unit testing function for toSnakeCase() function
func TestToSnakeCase(t *testing.T) {
	examples := []struct {
		str        string
		underscore string
	}{
		{"hi", "hi"},
		{"Hi", "hi"},
		{"HI", "hi"},
		{"HiWorld", "hi_world"},
	}
	for _, e := range examples {
		if toSnakeCase(e.str) != e.underscore {
			t.Errorf("toSnakeCase(%s) = %s != %s", e.str, toSnakeCase(e.str), e.underscore)
		}
	}
}

// TestJSONMarshal is a unit testing function for JSONMarshal() function
func TestJSONMarshal(t *testing.T) {
	examples := []struct {
		obj     interface{}
		safe    bool
		err     error
		rawjson string
	}{
		{struct {
			X int
		}{
			X: 5,
		}, true, nil, "{\n \"X\": 5\n}"},
		{struct {
			X string
		}{
			X: "hi",
		}, true, nil, "{\n \"X\": \"hi\"\n}"},
		{struct {
			X string
			Y int
		}{
			X: "hi",
			Y: 5,
		}, true, nil, "{\n \"X\": \"hi\",\n \"Y\": 5\n}"},
		{map[string]interface{}{
			"X": "hi",
			"Y": 5,
		}, true, nil, "{\n \"X\": \"hi\",\n \"Y\": 5\n}"},
	}
	for _, e := range examples {
		rawjson, err := JSONMarshal(e.obj, e.safe)
		if string(rawjson) != e.rawjson || err != e.err {
			t.Errorf("JSONMarshal(%#v, %v) = %s, %v != %s, %v", e.obj, e.safe, rawjson, err, e.rawjson, e.err)
		}
	}
}

// TestReturnJSON is a unit testing function for ReturnJSON() function
func TestReturnJSON(t *testing.T) {
	examples := []struct {
		m   interface{}
		out string
	}{
		{map[string]interface{}{"ID": 1, "Name": "Test"}, `{
  "ID": 1,
  "Name": "Test"
}`},
		{math.NaN(), `{
  "error_msg": "unable to encode JSON. json: unsupported value: NaN",
  "status": "error"
}`},
	}

	r := httptest.NewRequest("GET", "/", nil)

	for _, e := range examples {
		w := httptest.NewRecorder()
		ReturnJSON(w, r, e.m)
		buf, err := ioutil.ReadAll(w.Body)
		if err != nil {
			t.Errorf("ReturnJSON returned an error. %s", err)
		}
		if string(buf) != e.out {
			t.Errorf("ReturnJSON returned invalid JSON. Expected %s, got %s", e.out, string(buf))
		}
	}
}
