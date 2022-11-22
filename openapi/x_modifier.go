package openapi

type XModifier struct {
	Modifier string `json:"modifier,omitempty"`
	In       string `json:"in,omitempty"`
	Summary  string `json:"summary,omitempty"`
}
