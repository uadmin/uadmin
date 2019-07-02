package uadmin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

// TestProcessDelete is a unit testing function for processDelete() function
func TestProcessDelete(t *testing.T) {
	// Setup
	w := httptest.NewRecorder()

	idList := []string{}

	for i := 0; i < 10; i++ {
		record := TestStruct{
			Name: fmt.Sprintf("test%d", i),
		}
		Save(&record)
		idList = append(idList, fmt.Sprint(record.ID))
	}

	examples := []struct {
		r     *http.Request
		count int
	}{
		{httptest.NewRequest("POST", "/", nil), 0},
		{httptest.NewRequest("POST", "/", nil), len(idList)},
	}

	examples[1].r.PostForm = url.Values{}
	examples[1].r.PostForm.Set("listID", strings.Join(idList, ","))

	user := User{}
	Get(&user, "id = ?", 1)

	for _, e := range examples {
		countBefore := Count(&TestStruct{}, "")
		processDelete("teststruct", w, e.r, nil, &user)
		countAfter := Count(TestStruct{}, "")
		if (countBefore - countAfter) != e.count {
			t.Errorf("Invalid number of deleted records by processDelete. Expected %d, Got %d", e.count, (countBefore - countAfter))
		}
	}

	// Clean up
	DeleteList(TestStruct{}, "")
}
