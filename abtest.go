package uadmin

type TestType int

func (TestType) Static() TestType {
	return 1
}

func (TestType) Model() TestType {
	return 2
}

type ModelList int
type FieldList int

type ABTest struct {
	Model
	Name       string   `uadmin:"required"`
	Type       TestType `uadmin:"required"`
	StaticPath string
	ModelName  ModelList
	Field      FieldList
	PrimaryKey int
	Active     bool
	Group      string
}

func (a *ABTest) Save() {
	Save(a)
	syncABTests()
}
