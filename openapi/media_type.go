package openapi

type MediaType struct {
	Schema   *SchemaObject       `json:"schema,omitempty"`
	Example  *Example            `json:"example,omitempty"`
	Examples map[string]Example  `json:"examples,omitempty"`
	Encoding map[string]Encoding `json:"encoding,omitempty"`
}
