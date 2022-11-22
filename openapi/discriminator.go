package openapi

type Discriminator struct {
	PropertyName string            `json:"propertyName,omitempty"`
	Mapping      map[string]string `json:"mapping,omitempty"`
}
