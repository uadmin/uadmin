package openapi

type SchemaObject struct {
	Ref           string                   `json:"$ref,omitempty"`
	Type          string                   `json:"type,omitempty"`
	Pattern       string                   `json:"pattern,omitempty"`
	Maximum       int                      `json:"maximum,omitempty"`
	Minimum       int                      `json:"minimum,omitempty"`
	Required      []string                 `json:"required,omitempty"`
	Title         string                   `json:"title,omitempty"`
	Description   string                   `json:"description,omitempty"`
	Default       string                   `json:"default,omitempty"`
	ReadOnly      bool                     `json:"ReadOnly,omitempty"`
	Examples      []Example                `json:"examples,omitempty"`
	Items         *SchemaObject            `json:"items,omitempty"`
	Properties    map[string]*SchemaObject `json:"properties,omitempty"`
	Discriminator *Discriminator           `json:"discriminator,omitempty"`
	XML           *XML                     `json:"xml,omitempty"`
	ExternalDocs  *ExternalDocs            `json:"externalDocs,omitempty"`
	Example       *Example                 `json:"example,omitempty"`
	Enum          []interface{}            `json:"enum,omitempty"`
	OneOf         []*SchemaObject          `json:"oneOf,omitempty"`
	Const         interface{}              `json:"const,omitempty"`
	XFilters      []XModifier              `json:"x-filter,omitempty"`
	XAggregator   []XModifier              `json:"x-aggregator,omitempty"`
}
