package uadmin

type BuilderField struct {
	Model
	Builder   Builder
	BuilderID uint
	Name      string
	DataType  DataType
}
