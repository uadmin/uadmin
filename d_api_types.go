package uadmin

import "strings"

// /modelname/delete/1/             Delete One
// /modelname/method/METHOD_NAME/1/ Run method on model where id=1
type DApiModelKeyVal struct {
	PathCommand     PathCommandType `param:"command"`
	PathCommandName string          `param:"command"`

	CommandName string

	// when processing /method/ DataCommand is equal to the METHOD_NAME
	DataCommand   string
	DataForMethod string
	// in case want to use an already available *Session
	session *Session
}

type PathCommandType uint64

const (
	Auth        PathCommandType = 0
	AllModels   PathCommandType = 1
	Help        PathCommandType = 2
	DataCommand PathCommandType = 3
)

func (s PathCommandType) String() string {
	switch s {
	case Auth:
		return "auth"
	case AllModels:
		return "$allmodels"
	case Help:
		return "help"
	case DataCommand:
		return "datacommand"
	}
	return "unknown"
}

var (
	PathCommandTypeMap = map[string]PathCommandType{
		"auth":        Auth,
		"$allmodels":  AllModels,
		"help":        Help,
		"datacommand": DataCommand,
	}
)

func ParseCommandString(str string) (PathCommandType, bool) {
	c, ok := PathCommandTypeMap[strings.ToLower(str)]
	return c, ok
}

const (
	ReadCommand   = "read"
	AddCommand    = "add"
	EditCommand   = "edit"
	DeleteCommand = "delete"
	SchemaCommand = "schema"
	MethodCommand = "method"
)

var (
	DataCommands = []string{"read", "add", "edit", "delete", "schema", "method"}
)
