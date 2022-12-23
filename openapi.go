package uadmin

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/uadmin/uadmin/openapi"
)

// CustomOpenAPI is a handler to be called to customize OpenAPI schema
// Use of OpenAPI schema generation is under development and should not be used in production
var CustomOpenAPI func(*openapi.Schema)

// CustomOpenAPIJSON is a handler to be called to customize OpenAPI schema JSON output
// Use of OpenAPI schema generation is under development and should not be used in production
var CustomOpenAPIJSON func([]byte) []byte

// GenerateOpenAPISchema generates API schema for dAPI that is compatible with OpenAPI 3.1.0
// Use of OpenAPI schema generation is under development and should not be used in production
func GenerateOpenAPISchema() {
	Trail(WARNING, "Use of OpenAPI schema generation is under development and should not be used in production")
	s := openapi.GenerateBaseSchema()

	// Customize schema
	s.Info.Title = SiteName + s.Info.Title

	// Add Models to /components/schema
	for _, v := range Schema {
		// Parse fields
		fields := map[string]*openapi.SchemaObject{}
		required := []string{}
		parameters := []openapi.Parameter{}
		writeParameters := map[string]*openapi.SchemaObject{}

		// Add tag to schema if it doesn't exist
		tag := "Other"
		if v.Category != "" {
			tag = v.Category

			// check if it exists
			tagExists := false
			for i := range s.Tags {
				if s.Tags[i].Name == tag {
					tagExists = true
					break
				}
			}

			// if it doesn't exist, add it
			if !tagExists {
				s.Tags = append(s.Tags, openapi.Tag{
					Name:        tag,
					Description: "CRUD APIs for " + tag + " models",
				})
			}
		}

		for i := range v.Fields {
			// Determine data type
			fields[v.Fields[i].Name] = func() *openapi.SchemaObject {
				switch v.Fields[i].Type {
				case cID:
					return &openapi.SchemaObject{
						Type: "integer",
					}
				case cSTRING:
					return &openapi.SchemaObject{
						Type: "string",
					}
				case cBOOL:
					return &openapi.SchemaObject{
						Type: "boolean",
					}
				case cCODE:
					return &openapi.SchemaObject{
						Type: "boolean",
					}
				case cDATE:
					return &openapi.SchemaObject{
						Type: "string",
					}
				case cEMAIL:
					return &openapi.SchemaObject{
						Type: "string",
					}
				case cFILE:
					return &openapi.SchemaObject{
						Type: "string",
					}
				case cIMAGE:
					return &openapi.SchemaObject{
						Type: "string",
					}
				case cFK:
					return &openapi.SchemaObject{
						AllOf: []*openapi.SchemaObject{
							{Ref: "#/components/schemas/" + v.Fields[i].TypeName},
							{},
						},
					}
				case cHTML:
					return &openapi.SchemaObject{
						Type: "string",
					}
				case cLINK:
					return &openapi.SchemaObject{
						Type: "string",
					}
				case cLIST:
					return &openapi.SchemaObject{
						Type: "integer",
						OneOf: func() []*openapi.SchemaObject {
							vals := make([]*openapi.SchemaObject, len(v.Fields[i].Choices))
							for j := range v.Fields[i].Choices {
								vals[j] = &openapi.SchemaObject{
									Type:  "integer",
									Title: v.Fields[i].Choices[j].V,
									Const: v.Fields[i].Choices[j].K,
								}
							}
							return vals
						}(),
					}
				case cM2M:
					return &openapi.SchemaObject{
						Type:  "array",
						Items: &openapi.SchemaObject{Ref: "#/components/schemas/" + v.Fields[i].TypeName},
					}
				case cMONEY:
					return &openapi.SchemaObject{
						Type: "number",
					}
				case cNUMBER:
					switch v.Fields[i].TypeName {
					case "float64":
						return &openapi.SchemaObject{
							Type: "number",
						}
					case "int":
						return &openapi.SchemaObject{
							Type: "integer",
						}
					default:
						return &openapi.SchemaObject{
							Type: "integer",
						}
					}
				case cMULTILINGUAL:
					return &openapi.SchemaObject{
						Type: "string",
					}
				case cPROGRESSBAR:
					switch v.Fields[i].TypeName {
					case "float64":
						return &openapi.SchemaObject{
							Type: "number",
						}
					case "int":
						return &openapi.SchemaObject{
							Type: "integer",
						}
					default:
						return &openapi.SchemaObject{
							Type: "integer",
						}
					}
				default:
					return &openapi.SchemaObject{
						Type: "string",
					}
				}

			}()

			// If the field is a foreign key, then add the ID field for it
			if v.Fields[i].Type == cFK {
				fields[v.Fields[i].Name+"ID"] = &openapi.SchemaObject{
					Type: "integer",
				}
			}

			// Set other schema properties
			if v.Fields[i].Type != cFK {
				fields[v.Fields[i].Name].Description = v.Fields[i].Help
				fields[v.Fields[i].Name].Default = v.Fields[i].DefaultValue
				fields[v.Fields[i].Name].Title = v.Fields[i].DisplayName
				fields[v.Fields[i].Name].ReadOnly = func() *bool {
					if val := v.Fields[i].ReadOnly != ""; val {
						return &val
					}
					return nil
				}()
				fields[v.Fields[i].Name].Pattern = v.Fields[i].Pattern
				fields[v.Fields[i].Name].Format = func() string {
					switch v.Fields[i].Type {
					case cDATE:
						return "date-time"
					case cPASSWORD:
						return "password"
					case cEMAIL:
						return "email"
					case cHTML:
						return "html"
					default:
						return ""
					}
				}()
				fields[v.Fields[i].Name].Deprecated = func() *bool {
					if v.Fields[i].Deprecated {
						return &v.Fields[i].Deprecated
					}
					return nil
				}()
				if val, ok := v.Fields[i].Max.(string); ok && val != "" {
					fields[v.Fields[i].Name].Maximum, _ = strconv.Atoi(val)
				}
				if val, ok := v.Fields[i].Min.(string); ok && val != "" {
					fields[v.Fields[i].Name].Minimum, _ = strconv.Atoi(val)
				}
			} else {
				fields[v.Fields[i].Name].AllOf[1].Description = v.Fields[i].Help
				fields[v.Fields[i].Name].AllOf[1].Default = v.Fields[i].DefaultValue
				fields[v.Fields[i].Name].AllOf[1].Title = v.Fields[i].DisplayName
				fields[v.Fields[i].Name].ReadOnly = func() *bool {
					if val := v.Fields[i].ReadOnly != ""; val {
						return &val
					}
					return nil
				}()
				fields[v.Fields[i].Name].Deprecated = func() *bool {
					if v.Fields[i].Deprecated {
						return &v.Fields[i].Deprecated
					}
					return nil
				}()
			}

			// Add parameters
			// skip method fields
			if v.Fields[i].IsMethod {
				continue
			}
			parameters = append(parameters, func() openapi.Parameter {
				if v.Fields[i].Type == cID {
					return openapi.Parameter{
						Ref: "#/components/parameters/QueryID",
					}
				}
				return openapi.Parameter{
					Name: func() string {
						if v.Fields[i].Type == cFK {
							return v.Fields[i].ColumnName + "_id"
						} else {
							return v.Fields[i].ColumnName
						}
					}(),
					In:          "query",
					Description: "Query for " + v.Fields[i].DisplayName,
					Schema: func() *openapi.SchemaObject {
						switch v.Fields[i].Type {
						case cSTRING:
							fallthrough
						case cCODE:
							fallthrough
						case cEMAIL:
							fallthrough
						case cFILE:
							fallthrough
						case cIMAGE:
							fallthrough
						case cHTML:
							fallthrough
						case cLINK:
							fallthrough
						case cMULTILINGUAL:
							fallthrough
						case cPASSWORD:
							return &openapi.SchemaObject{
								Ref: "#/components/schemas/String",
							}
						case cFK:
							return &openapi.SchemaObject{
								Ref: "#/components/schemas/Integer",
							}
						case cLIST:
							return &openapi.SchemaObject{
								Type: "integer",
								OneOf: func() []*openapi.SchemaObject {
									vals := make([]*openapi.SchemaObject, len(v.Fields[i].Choices))
									for j := range v.Fields[i].Choices {
										vals[j] = &openapi.SchemaObject{
											Type:  "integer",
											Title: v.Fields[i].Choices[j].V,
											Const: v.Fields[i].Choices[j].K,
										}
									}
									return vals
								}(),
							}
						case cMONEY:
							fallthrough
						case cNUMBER:
							fallthrough
						case cPROGRESSBAR:
							return &openapi.SchemaObject{
								Ref: "#/components/schemas/Number",
							}
						case cBOOL:
							return &openapi.SchemaObject{
								Ref: "#/components/schemas/Boolean",
							}
						case cDATE:
							return &openapi.SchemaObject{
								Ref: "#/components/schemas/DateTime",
							}
						default:
							return &openapi.SchemaObject{Ref: "#/components/schemas/String"}
						}
					}(),
				}
			}(),
			)

			if v.Fields[i].Type == cID {
				continue
			}

			writeParameterName := func() string {
				if v.Fields[i].Type == cFK {
					return "_" + v.Fields[i].ColumnName + "_id"
				} else {
					return "_" + v.Fields[i].ColumnName
				}
			}()
			writeParameters[writeParameterName] = func() *openapi.SchemaObject {
				switch v.Fields[i].Type {
				case cSTRING:
					fallthrough
				case cCODE:
					fallthrough
				case cEMAIL:
					fallthrough
				case cHTML:
					fallthrough
				case cLINK:
					fallthrough
				case cMULTILINGUAL:
					fallthrough
				case cPASSWORD:
					return &openapi.SchemaObject{
						Type:        "string",
						Description: v.Fields[i].Help,
					}
				case cFILE:
					fallthrough
				case cIMAGE:
					return &openapi.SchemaObject{
						Type:        "string",
						Format:      "binary",
						Description: v.Fields[i].Help,
					}
				case cFK:
					return &openapi.SchemaObject{
						Type:        "integer",
						Description: "Foreign key to " + v.Fields[i].TypeName + ". " + v.Fields[i].Help,
					}
				case cLIST:
					fallthrough
				case cMONEY:
					return &openapi.SchemaObject{
						Type:        "integer",
						Description: v.Fields[i].Help,
					}
				case cNUMBER:
					fallthrough
				case cPROGRESSBAR:
					return &openapi.SchemaObject{
						Type:        "number",
						Description: v.Fields[i].Help,
					}
				case cBOOL:
					return &openapi.SchemaObject{
						Type:        "string",
						Enum:        []interface{}{"", "0", "1"},
						Description: v.Fields[i].Help,
					}
				case cDATE:
					return &openapi.SchemaObject{
						Type:        "string",
						Description: v.Fields[i].Help,
					}
				default:
					return &openapi.SchemaObject{
						Type:        "string",
						Description: v.Fields[i].Help,
					}
				}
			}()

			// Add required fields
			if v.Fields[i].Required {
				required = append(required, v.Fields[i].Name)
			}
		}

		// Add dAPI paths
		// Read one
		s.Paths[fmt.Sprintf("/api/d/%s/{id}", v.ModelName)] = openapi.Path{
			Summary:     "Single record operations for " + v.DisplayName,
			Description: "Single record operations for " + v.DisplayName,
			Get: &openapi.Operation{
				Tags:        []string{tag},
				Summary:     "Read one " + v.DisplayName + " record",
				Description: "Read one " + v.DisplayName + " record",
				Responses: map[string]openapi.Response{
					"200": {
						Description: v.DisplayName + " record",
						Content: map[string]openapi.MediaType{
							"application/json": {
								Schema: &openapi.SchemaObject{
									Type: "object",
									Properties: map[string]*openapi.SchemaObject{
										"result": {Ref: "#/components/schemas/" + v.Name},
										"status": {Type: "string"},
									},
								},
							},
						},
					},
				},
				Parameters: []openapi.Parameter{
					{
						Ref: "#/components/parameters/PathID",
					},
					{
						Ref: "#/components/parameters/deleted",
					},
					{
						Ref: "#/components/parameters/m2m",
					},
					{
						Ref: "#/components/parameters/preload",
					},
					{
						Ref: "#/components/parameters/stat",
					},
				},
			},
			Put: &openapi.Operation{
				Tags:        []string{tag},
				Summary:     "Edit one " + v.DisplayName + " record",
				Description: "Edit one " + v.DisplayName + " record",
				RequestBody: &openapi.RequestBody{
					Content: map[string]openapi.MediaType{
						"multipart/form-data": {
							Schema: &openapi.SchemaObject{
								Type:       "object",
								Properties: writeParameters,
							},
						},
					},
				},
				Responses: map[string]openapi.Response{
					"200": {
						Description: v.DisplayName + " record edited",
						Content: map[string]openapi.MediaType{
							"application/json": {
								Schema: &openapi.SchemaObject{
									Type: "object",
									Properties: map[string]*openapi.SchemaObject{
										"rows_count": {Type: "integer"},
										"status":     {Type: "string"},
									},
								},
							},
						},
					},
				},
				Parameters: []openapi.Parameter{
					{
						Ref: "#/components/parameters/PathID",
					},
					{
						Ref: "#/components/parameters/CSRF",
					},
					{
						Ref: "#/components/parameters/stat",
					},
				},
			},
			Delete: &openapi.Operation{
				Tags:        []string{tag},
				Summary:     "Delete one " + v.DisplayName + " record",
				Description: "Delete one " + v.DisplayName + " record",
				Responses: map[string]openapi.Response{
					"200": {
						Description: v.DisplayName + " record deleted",
						Content: map[string]openapi.MediaType{
							"application/json": {
								Schema: &openapi.SchemaObject{
									Type: "object",
									Properties: map[string]*openapi.SchemaObject{
										"rows_count": {Type: "integer"},
										"status":     {Type: "string"},
									},
								},
							},
						},
					},
				},
				Parameters: []openapi.Parameter{
					{
						Ref: "#/components/parameters/PathID",
					},
					{
						Ref: "#/components/parameters/CSRF",
					},
					{
						Ref: "#/components/parameters/stat",
					},
				},
			},
		}
		// Read Many
		s.Paths[fmt.Sprintf("/api/d/%s", v.ModelName)] = openapi.Path{
			Summary:     "Add one and multi-record operations for " + v.DisplayName,
			Description: "Add one and multi-record operations for " + v.DisplayName,
			Get: &openapi.Operation{
				Tags:        []string{tag},
				Summary:     "Read many " + v.DisplayName + " records",
				Description: "Read many " + v.DisplayName + " records",
				Responses: map[string]openapi.Response{
					"200": {
						Description: v.DisplayName + " records",
						Content: map[string]openapi.MediaType{
							"application/json": {
								Schema: &openapi.SchemaObject{
									Type: "object",
									Properties: map[string]*openapi.SchemaObject{
										"result": {
											Type:  "array",
											Items: &openapi.SchemaObject{Ref: "#/components/schemas/" + v.Name},
										},
										"status": {Type: "string"},
									},
								},
							},
						},
					},
				},
				Parameters: append(parameters, []openapi.Parameter{
					{
						Ref: "#/components/parameters/limit",
					},
					{
						Ref: "#/components/parameters/offset",
					},
					{
						Ref: "#/components/parameters/order",
					},
					{
						Ref: "#/components/parameters/fields",
					},
					{
						Ref: "#/components/parameters/groupBy",
					},
					{
						Ref: "#/components/parameters/deleted",
					},
					{
						Ref: "#/components/parameters/join",
					},
					{
						Ref: "#/components/parameters/m2m",
					},
					{
						Ref: "#/components/parameters/q",
					},
					{
						Ref: "#/components/parameters/stat",
					},
					{
						Ref: "#/components/parameters/or",
					},
				}...),
			},
			Post: &openapi.Operation{
				Tags:        []string{tag},
				Summary:     "Add one " + v.DisplayName + " record",
				Description: "Add one " + v.DisplayName + " record",
				RequestBody: &openapi.RequestBody{
					Content: map[string]openapi.MediaType{
						"multipart/form-data": {
							Schema: &openapi.SchemaObject{
								Type:       "object",
								Properties: writeParameters,
							},
						},
					},
				},
				Responses: map[string]openapi.Response{
					"200": {
						Description: v.DisplayName + " record added",
						Content: map[string]openapi.MediaType{
							"application/json": {
								Schema: &openapi.SchemaObject{
									Type: "object",
									Properties: map[string]*openapi.SchemaObject{
										"id": {
											Type:  "array",
											Items: &openapi.SchemaObject{Type: "integer"},
										},
										"rows_count": {Type: "integer"},
										"status":     {Type: "string"},
									},
								},
							},
						},
					},
				},
				Parameters: []openapi.Parameter{
					{
						Ref: "#/components/parameters/CSRF",
					},
					{
						Ref: "#/components/parameters/stat",
					},
				},
			},
			Put: &openapi.Operation{
				Tags:        []string{tag},
				Summary:     "Edit many " + v.DisplayName + " records",
				Description: "Edit many " + v.DisplayName + " records",
				RequestBody: &openapi.RequestBody{
					Content: map[string]openapi.MediaType{
						"multipart/form-data": {
							Schema: &openapi.SchemaObject{
								Type:       "object",
								Properties: writeParameters,
							},
						},
					},
				},
				Responses: map[string]openapi.Response{
					"200": {
						Description: v.DisplayName + " records edited",
						Content: map[string]openapi.MediaType{
							"application/json": {
								Schema: &openapi.SchemaObject{
									Type: "object",
									Properties: map[string]*openapi.SchemaObject{
										"rows_count": {Type: "integer"},
										"status":     {Type: "string"},
									},
								},
							},
						},
					},
				},
				Parameters: append([]openapi.Parameter{
					{
						Ref: "#/components/parameters/CSRF",
					},
					{
						Ref: "#/components/parameters/stat",
					},
				}, parameters...),
			},
			Delete: &openapi.Operation{
				Tags:        []string{tag},
				Summary:     "Delete many " + v.DisplayName + " records",
				Description: "Delete many " + v.DisplayName + " records",
				Responses: map[string]openapi.Response{
					"200": {
						Description: v.DisplayName + " records deleted",
						Content: map[string]openapi.MediaType{
							"application/json": {
								Schema: &openapi.SchemaObject{
									Type: "object",
									Properties: map[string]*openapi.SchemaObject{
										"rows_count": {Type: "integer"},
										"status":     {Type: "string"},
									},
								},
							},
						},
					},
				},
				Parameters: append([]openapi.Parameter{
					{
						Ref: "#/components/parameters/CSRF",
					},
					{
						Ref: "#/components/parameters/stat",
					},
				}, parameters...),
			},
		}
		s.Components.Schemas[v.Name] = openapi.SchemaObject{
			Type:       "object",
			Properties: fields,
			Required:   required,
		}
	}

	// run custom OpenAPI handler
	if CustomOpenAPI != nil {
		CustomOpenAPI(s)
	}

	buf := getOpenAPIJSON(s)
	os.WriteFile("./openapi.json", buf, 0644)
}

func getOpenAPIJSON(s *openapi.Schema) []byte {
	buf, err := json.MarshalIndent(*s, "", "  ")
	if err != nil {
		return nil
	}
	if CustomOpenAPIJSON != nil {
		buf = CustomOpenAPIJSON(buf)
	}
	return buf
}
