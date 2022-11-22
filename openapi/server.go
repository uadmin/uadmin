package openapi

type Server struct {
	URL         string                           `json:"url,omitempty"`
	Description string                           `json:"description,omitempty"`
	Variables   map[string]OpenAPIServerVariable `json:"variables,omitempty"`
}
