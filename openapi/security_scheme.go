package openapi

type SecurityScheme struct {
	Ref              string      `json:"$ref,omitempty"`
	Type             string      `json:"type,omitempty"`
	Description      string      `json:"description,omitempty"`
	Name             string      `json:"name,omitempty"`
	In               string      `json:"in,omitempty"`
	Scheme           string      `json:"scheme,omitempty"`
	BearerFormat     string      `json:"bearerFormat,omitempty"`
	Flows            *OAuthFlows `json:"flows,omitempty"`
	OpenIdConnectUrl string      `json:"openIdConnectUrl,omitempty"`
}
