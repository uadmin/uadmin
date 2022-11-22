package openapi

type Schema struct {
	OpenAPI           string                `json:"openapi,omitempty"`
	Info              *SchemaInfo           `json:"info,omitempty"`
	JSONSchemaDialect string                `json:"jsonSchemaDialect,omitempty"`
	Servers           []Server              `json:"servers,omitempty"`
	Paths             map[string]Path       `json:"paths,omitempty"`
	Webhooks          map[string]Path       `json:"webhooks,omitempty"`
	Components        *Components           `json:"components,omitempty"`
	Security          []SecurityRequirement `json:"security,omitempty"`
	Tags              []Tag                 `json:"tags,omitempty"`
	ExternalDocs      *ExternalDocs         `json:"externalDocs,omitempty"`
}
