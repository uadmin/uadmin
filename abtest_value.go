package uadmin

import (
	"fmt"
)

type ABTestValue struct {
	Model
	ABTest      ABTest
	ABTestID    uint
	Value       string
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

func (ABTestValue) HideInDashboard() bool {
	return true
}
