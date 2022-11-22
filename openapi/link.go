package openapi

type Link struct {
	Ref          string                 `json:"$ref,omitempty"`
	OperationRef string                 `json:"operationRef,omitempty"`
	OperationId  string                 `json:"operationId,omitempty"`
	Parameters   map[string]interface{} `json:"parameters,omitempty"`
	RequestBody  map[string]interface{} `json:"requestBody,omitempty"`
	Description  string                 `json:"description,omitempty"`
	Server       *Server                `json:"server,omitempty"`
}
