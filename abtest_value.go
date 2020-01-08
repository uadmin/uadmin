package uadmin

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type ABTestValue struct {
	Model
	ABTest      ABTest
	ABTestID    uint
	Value       string `uadmin:"list_exclude"`
	Active      bool
	Impressions int
	Clicks      int
}

func (a *ABTestValue) String() string {
	return a.Value
}

func (a *ABTestValue) ClickThroughRate() float64 {
	if a.Impressions == 0 {
		return 0.0
	}
	return float64(a.Clicks) / float64(a.Impressions) * 100
}

func (a ABTestValue) ClickThroughRate__Form__List() string {
	return fmt.Sprintf("<b>%.2f%%</b>", a.ClickThroughRate())
}

func (a ABTestValue) Preview__Form__List() string {
	// Check if the value is a path to a file
	if strings.HasPrefix(a.Value, "/") {
		// Check the file type
		// Image
		if strings.HasSuffix(a.Value, "png") || strings.HasSuffix(a.Value, "jpg") || strings.HasSuffix(a.Value, "gif") || strings.HasSuffix(a.Value, "jpeg") {
			return fmt.Sprintf(`<img src="%s" style="width:256px">`, a.Value)
		}
		// CSS/JS
		if strings.HasSuffix(a.Value, "css") || strings.HasSuffix(a.Value, "js") {
			buf, _ := ioutil.ReadFile("." + a.Value)
			return fmt.Sprintf(`<pre style="width:256px">%s\n%s</pre>`, a.Value, string(buf))
		}
	}
	return a.Value
}

func (ABTestValue) HideInDashboard() bool {
	return true
}
