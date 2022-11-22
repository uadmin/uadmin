package uadmin

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/uadmin/uadmin/openapi"
)

var CustomOpenAPI func(*openapi.Schema)
var CustomOpenAPIJSON func([]byte) []byte

func GenerateOpenAPISchema() {
	s := openapi.GenerateBaseSchema()

	// Customize schema
	s.Info.Title = SiteName + s.Info.Title

	// Add Models to /components/schema
	for _, v := range Schema {
		// Parse fields
		fields := map[string]*openapi.SchemaObject{}
		required := []string{}
		parameters := []openapi.Parameter{}
		writeParameters := []openapi.Parameter{}
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
						Type:  "object",
						Items: &openapi.SchemaObject{Ref: "#/components/schemas/" + v.Fields[i].TypeName},
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
						Type: "string",
						Enum: func() []interface{} {
							vals := make([]interface{}, len(v.Fields[i].Choices))
							for j := range v.Fields[i].Choices {
								vals[j] = v.Fields[i].Choices[j].V
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
				}
				// If unknown, use string
				return &openapi.SchemaObject{
					Type: "string",
				}
			}()

			// Set other schema properties
			fields[v.Fields[i].Name].Description = v.Fields[i].Help
			fields[v.Fields[i].Name].Default = v.Fields[i].DefaultValue
			fields[v.Fields[i].Name].Title = v.Fields[i].DisplayName
			if val, ok := v.Fields[i].Max.(string); ok && val != "" {
				fields[v.Fields[i].Name].Maximum, _ = strconv.Atoi(val)
			}
			if val, ok := v.Fields[i].Min.(string); ok && val != "" {
				fields[v.Fields[i].Name].Minimum, _ = strconv.Atoi(val)
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
						case cCODE:
						case cEMAIL:
						case cFILE:
						case cIMAGE:
						case cHTML:
						case cLINK:
						case cMULTILINGUAL:
						case cPASSWORD:
							return &openapi.SchemaObject{
								Ref: "#/components/schemas/String",
							}
						case cID:
						case cFK:
						case cLIST:
						case cMONEY:
							return &openapi.SchemaObject{
								Ref: "#/components/schemas/Integer",
							}
						case cNUMBER:
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
						}
						return &openapi.SchemaObject{
							Ref: "#/components/schemas/String",
						}
					}(),
				}
			}(),
			)

			if v.Fields[i].Type == cID {
				continue
			}

			writeParameters = append(writeParameters, func() openapi.Parameter {
				return openapi.Parameter{
					Name: func() string {
						if v.Fields[i].Type == cFK {
							return "_" + v.Fields[i].ColumnName + "_id"
						} else {
							return "_" + v.Fields[i].ColumnName
						}
					}(),
					In:          "query",
					Description: "Set value for " + v.Fields[i].DisplayName,
					Schema: func() *openapi.SchemaObject {
						switch v.Fields[i].Type {
						case cSTRING:
						case cCODE:
						case cEMAIL:
						case cFILE:
						case cIMAGE:
						case cHTML:
						case cLINK:
						case cMULTILINGUAL:
						case cPASSWORD:
							return &openapi.SchemaObject{
								Ref: "#/components/schemas/String",
							}
						case cID:
						case cFK:
						case cLIST:
						case cMONEY:
							return &openapi.SchemaObject{
								Ref: "#/components/schemas/Integer",
							}
						case cNUMBER:
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
						}
						return &openapi.SchemaObject{
							Ref: "#/components/schemas/String",
						}
					}(),
				}
			}(),
			)

			// Add required fields
			if v.Fields[i].Required {
				required = append(required, v.Fields[i].Name)
			}
		}

		// Add dAPI paths
		// Read one
		s.Paths[fmt.Sprintf("/api/d/%s/read/{id}", v.ModelName)] = openapi.Path{
			Summary:     "Read one " + v.DisplayName,
			Description: "Read one " + v.DisplayName,
			Get: &openapi.Operation{
				Tags: []string{v.Name, func() string {
					if v.Category != "" {
						return v.Category
					} else {
						return "Other"
					}
				}()},
				Responses: map[string]openapi.Response{
					"200": {
						Description: v.DisplayName + " record",
						Content: map[string]openapi.MediaType{
							"application/json": {
								Schema: &openapi.SchemaObject{
									Ref: "#/components/schemas/" + v.Name,
								},
							},
						},
					},
				},
			},
			Post: &openapi.Operation{
				Tags: []string{v.Name, func() string {
					if v.Category != "" {
						return v.Category
					} else {
						return "Other"
					}
				}()},
				Responses: map[string]openapi.Response{
					"200": {
						Description: v.DisplayName + " record",
						Content: map[string]openapi.MediaType{
							"application/json": {
								Schema: &openapi.SchemaObject{
									Ref: "#/components/schemas/" + v.Name,
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
		}
		// Read Many
		s.Paths[fmt.Sprintf("/api/d/%s/read", v.ModelName)] = openapi.Path{
			Summary:     "Read one " + v.DisplayName,
			Description: "Read one " + v.DisplayName,
			Get: &openapi.Operation{
				Tags: []string{v.Name, func() string {
					if v.Category != "" {
						return v.Category
					} else {
						return "Other"
					}
				}()},
				Responses: map[string]openapi.Response{
					"200": {
						Description: v.DisplayName + " record",
						Content: map[string]openapi.MediaType{
							"application/json": {
								Schema: &openapi.SchemaObject{
									Ref: "#/components/schemas/" + v.Name,
								},
							},
						},
					},
				},
			},
			Post: &openapi.Operation{
				Tags: []string{v.Name, func() string {
					if v.Category != "" {
						return v.Category
					} else {
						return "Other"
					}
				}()},
				Responses: map[string]openapi.Response{
					"200": {
						Description: v.DisplayName + " record",
						Content: map[string]openapi.MediaType{
							"application/json": {
								Schema: &openapi.SchemaObject{
									Ref: "#/components/schemas/" + v.Name,
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
		}
		// Add One
		s.Paths[fmt.Sprintf("/api/d/%s/add", v.ModelName)] = openapi.Path{
			Summary:     "Add one " + v.DisplayName,
			Description: "Add one " + v.DisplayName,
			Get: &openapi.Operation{
				Tags: []string{v.Name, func() string {
					if v.Category != "" {
						return v.Category
					} else {
						return "Other"
					}
				}()},
				Responses: map[string]openapi.Response{
					"200": {
						Description: v.DisplayName + " record",
						Content: map[string]openapi.MediaType{
							"application/json": {
								Schema: &openapi.SchemaObject{
									Ref: "#/components/schemas/" + v.Name,
								},
							},
						},
					},
				},
			},
			Post: &openapi.Operation{
				Tags: []string{v.Name, func() string {
					if v.Category != "" {
						return v.Category
					} else {
						return "Other"
					}
				}()},
				Responses: map[string]openapi.Response{
					"200": {
						Description: v.DisplayName + " record",
						Content: map[string]openapi.MediaType{
							"application/json": {
								Schema: &openapi.SchemaObject{
									Ref: "#/components/schemas/" + v.Name,
								},
							},
						},
					},
				},
			},
			Parameters: append(writeParameters, []openapi.Parameter{
				{
					Ref: "#/components/parameters/CSRF",
				},
				{
					Ref: "#/components/parameters/stat",
				},
			}...),
		}
		// Add One
		s.Paths[fmt.Sprintf("/api/d/%s/add", v.ModelName)] = openapi.Path{
			Summary:     "Add one " + v.DisplayName,
			Description: "Add one " + v.DisplayName,
			Get: &openapi.Operation{
				Tags: []string{v.Name, func() string {
					if v.Category != "" {
						return v.Category
					} else {
						return "Other"
					}
				}()},
				Responses: map[string]openapi.Response{
					"200": {
						Description: v.DisplayName + " record",
						Content: map[string]openapi.MediaType{
							"application/json": {
								Schema: &openapi.SchemaObject{
									Ref: "#/components/schemas/" + v.Name,
								},
							},
						},
					},
				},
			},
			Post: &openapi.Operation{
				Tags: []string{v.Name, func() string {
					if v.Category != "" {
						return v.Category
					} else {
						return "Other"
					}
				}()},
				Responses: map[string]openapi.Response{
					"200": {
						Description: v.DisplayName + " record",
						Content: map[string]openapi.MediaType{
							"application/json": {
								Schema: &openapi.SchemaObject{
									Ref: "#/components/schemas/" + v.Name,
								},
							},
						},
					},
				},
			},
			Parameters: append(writeParameters, []openapi.Parameter{
				{
					Ref: "#/components/parameters/CSRF",
				},
				{
					Ref: "#/components/parameters/stat",
				},
			}...),
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
	buf, err := json.Marshal(*s)
	if err != nil {
		return nil
	}
	if CustomOpenAPIJSON != nil {
		buf = CustomOpenAPIJSON(buf)
	}
	return buf
}
