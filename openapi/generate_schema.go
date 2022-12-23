package openapi

func GenerateBaseSchema() *Schema {
	s := &Schema{
		OpenAPI: "3.0.3",
		Info: &SchemaInfo{
			Title:       " API documentation",
			Description: "API documentation",
			Version:     "1.0.0",
		},
		Tags: []Tag{
			{
				Name:        "Auth",
				Description: "Authentication API",
			},
			{
				Name:        "System",
				Description: "CRUD APIs for uAdmin core models",
			},
			{
				Name:        "Other",
				Description: "CRUD APIs for models with no category",
			},
		},
		Components: &Components{
			Schemas: map[string]SchemaObject{
				"Integer": {
					Type: "integer",
					XFilters: []XModifier{
						{Modifier: "__gt", In: "suffix", Summary: "Greater than"},
						{Modifier: "__gte", In: "suffix", Summary: "Greater than or equal to"},
						{Modifier: "__lt", In: "suffix", Summary: "Less than"},
						{Modifier: "__lte", In: "suffix", Summary: "Less than or equal to"},
						{Modifier: "__in", In: "suffix", Summary: "Find a value matching any of these values"},
						{Modifier: "__between", In: "suffix", Summary: "Selects values within a given range"},
						{Modifier: "!", In: "prefix", Summary: "Negates operator"},
					},
					XAggregator: []XModifier{
						{Modifier: "__sum", In: "suffix", Summary: "Returns the total sum of a numeric field"},
						{Modifier: "__avg", In: "suffix", Summary: "Returns the average value of a numeric field"},
						{Modifier: "__min", In: "suffix", Summary: "Returns the smallest value of a numeric field"},
						{Modifier: "__max", In: "suffix", Summary: "Returns the largest value of a numeric field"},
						{Modifier: "__count", In: "suffix", Summary: "Returns the number of rows"},
					},
				},
				"Number": {
					Type: "number",
					XFilters: []XModifier{
						{Modifier: "__gt", In: "suffix", Summary: "Greater than"},
						{Modifier: "__gte", In: "suffix", Summary: "Greater than or equal to"},
						{Modifier: "__lt", In: "suffix", Summary: "Less than"},
						{Modifier: "__lte", In: "suffix", Summary: "Less than or equal to"},
						{Modifier: "__in", In: "suffix", Summary: "Find a value matching any of these values"},
						{Modifier: "__between", In: "suffix", Summary: "Selects values within a given range"},
						{Modifier: "!", In: "prefix", Summary: "Negates operator"},
					},
					XAggregator: []XModifier{
						{Modifier: "__sum", In: "suffix", Summary: "Returns the total sum of a numeric field"},
						{Modifier: "__avg", In: "suffix", Summary: "Returns the average value of a numeric field"},
						{Modifier: "__min", In: "suffix", Summary: "Returns the smallest value of a numeric field"},
						{Modifier: "__max", In: "suffix", Summary: "Returns the largest value of a numeric field"},
						{Modifier: "__count", In: "suffix", Summary: "Returns the number of rows"},
					},
				},
				"String": {
					Type: "string",
					XFilters: []XModifier{
						{Modifier: "__contains", In: "suffix", Summary: "Search for string values that contains"},
						{Modifier: "__startswith", In: "suffix", Summary: "Search for string values that starts with a given substring"},
						{Modifier: "__endswith", In: "suffix", Summary: "Search for string values that ends with a given substring"},
						{Modifier: "__re", In: "suffix", Summary: "Search for string values that matches regular expression"},
						{Modifier: "__icontains", In: "suffix", Summary: "Search for string values that contains"},
						{Modifier: "__istartswith", In: "suffix", Summary: "Search for string values that starts with a given substring"},
						{Modifier: "__iendswith", In: "", Summary: "Search for string values that ends with a given substring"},
						{Modifier: "__in", In: "", Summary: "Find a value matching any of these values"},
						{Modifier: "!", In: "prefix", Summary: "Negates operator"},
					},
					XAggregator: []XModifier{
						{Modifier: "__count", In: "suffix", Summary: "Returns the number of rows"},
					},
				},
				"DateTime": {
					Type: "string",
					XFilters: []XModifier{
						{Modifier: "__contains", In: "suffix", Summary: "Search for string values that contains"},
						{Modifier: "__startswith", In: "suffix", Summary: "Search for string values that starts with a given substring"},
						{Modifier: "__endswith", In: "suffix", Summary: "Search for string values that ends with a given substring"},
						{Modifier: "__re", In: "suffix", Summary: "Search for string values that matches regular expression"},
						{Modifier: "__icontains", In: "suffix", Summary: "Search for string values that contains"},
						{Modifier: "__istartswith", In: "suffix", Summary: "Search for string values that starts with a given substring"},
						{Modifier: "__iendswith", In: "", Summary: "Search for string values that ends with a given substring"},
						{Modifier: "__in", In: "", Summary: "Find a value matching any of these values"},
						{Modifier: "!", In: "prefix", Summary: "Negates operator"},
					},
					XAggregator: []XModifier{
						{Modifier: "__count", In: "suffix", Summary: "Returns the number of rows"},
						{Modifier: "__date", In: "suffix", Summary: "Returns DATE() of the field"},
						{Modifier: "__year", In: "suffix", Summary: "Returns YEAR() of the field"},
						{Modifier: "__month", In: "suffix", Summary: "Returns MONTH() of the field"},
						{Modifier: "__day", In: "suffix", Summary: "Returns DAY() of the field"},
					},
				},
				"Boolean": {
					Type: "boolean",
					XAggregator: []XModifier{
						{Modifier: "__count", In: "suffix", Summary: "Returns the number of rows"},
					},
				},
				"GeneralError": {
					Type: "object",
					Properties: map[string]*SchemaObject{
						"status": {
							Type: "string",
						},
						"err_msg": {
							Type: "string",
						},
					},
				},
			},
			Parameters: map[string]Parameter{
				"PathID": {
					Name:        "id",
					In:          "path",
					Description: "Primary key of the record",
					Required:    true,
					Schema: &SchemaObject{
						Type: "integer",
					},
				},
				"QueryID": {
					Name:        "id",
					In:          "query",
					Description: "Primary key of the record",
					Required:    false,
					Schema: &SchemaObject{
						Ref: "#/components/schemas/Integer",
					},
				},
				"CSRF": {
					Name:        "X-CSRF-TOKEN",
					In:          "header",
					Description: "Token for CSRF protection which should be set to the session token",
					Required:    true,
					Schema: &SchemaObject{
						Type: "string",
					},
				},
				"limit": {
					Name:          "$limit",
					In:            "query",
					Description:   "Maximum number of records to return",
					Required:      false,
					AllowReserved: true,
					Schema: &SchemaObject{
						Type: "integer",
					},
				},
				"offset": {
					Name:          "$offset",
					In:            "query",
					Description:   "Starting point to read in the list of records",
					Required:      false,
					AllowReserved: true,
					Schema: &SchemaObject{
						Type: "integer",
					},
				},
				"order": {
					Name:          "$order",
					In:            "query",
					Description:   "Sort the results. Use '-' for descending order and comma for more field",
					Required:      false,
					AllowReserved: true,
					Schema: &SchemaObject{
						Type:    "string",
						Default: "",
					},
					Examples: map[string]Example{
						"multiColumn": {
							Summary: "An example multi-column sorting with ascending and descending",
							Value:   "$order=id,-name",
						},
					},
				},
				"fields": {
					Name:          "$f",
					In:            "query",
					Description:   "Selecting fields to return in results",
					Required:      false,
					AllowReserved: true,
					Schema: &SchemaObject{
						Type:    "string",
						Default: "",
					},
					Examples: map[string]Example{
						"multiColumn": {
							Summary: "An example multi-column selection",
							Value:   "$f=id,name",
						},
						"aggColumn": {
							Summary: "An example multi-column selection with aggregation function",
							Value:   "$f=id__count,score__sum",
						},
						"joinTable": {
							Summary: "An example of multi-column selection from a different table using a join (see $join)",
							Value:   "$f=username,user_groups.group_name",
						},
					},
				},
				"groupBy": {
					Name:          "$groupby",
					In:            "query",
					Description:   "Groups rows that have the same values into summary rows",
					Required:      false,
					AllowReserved: true,
					Schema: &SchemaObject{
						Type:    "string",
						Default: "",
					},
					Examples: map[string]Example{
						"simple": {
							Summary: "An example of grouping results based on category",
							Value:   "$groupby=category_id",
						},
						"agg": {
							Summary: "An example of grouping results based on year and month",
							Value:   "$groupby=date__year,date__month",
						},
					},
				},
				"deleted": {
					Name:            "$deleted",
					In:              "query",
					Description:     "Returns results including deleted records",
					Required:        false,
					AllowReserved:   true,
					AllowEmptyValue: true,
					Schema: &SchemaObject{
						Type:    "string",
						Default: "",
						Enum:    []interface{}{"", "true"},
					},
					Examples: map[string]Example{
						"getDeleted": {
							Summary: "An example of a query that returns deleted records",
							Value:   "$deleted=true",
						},
					},
				},
				"join": {
					Name:          "$join",
					In:            "query",
					Description:   "Joins results from another model based on a foreign key",
					Required:      false,
					AllowReserved: true,
					Schema: &SchemaObject{
						Type:    "string",
						Default: "",
					},
					Examples: map[string]Example{
						"getGroupName": {
							Summary: "An example of a query with a left join from users to user_groups",
							Value:   "$join=user_groups__left__user_group_id",
						},
						"getGroupNameInner": {
							Summary: "An example of a query with a inner join from users to user_groups",
							Value:   "$join=user_groups__user_group_id",
						},
					},
				},
				"m2m": {
					Name:          "$m2m",
					In:            "query",
					Description:   "Returns results from M2M fields",
					Required:      false,
					AllowReserved: true,
					Schema: &SchemaObject{
						Type:        "string",
						Enum:        []interface{}{"", "0", "fill", "id"},
						Default:     "",
						Description: "0=don't get, 1/fill=get full records, id=get ids only",
					},
					Examples: map[string]Example{
						"fillAll": {
							Summary: "An example of a query that fills all m2m records",
							Value:   "$m2m=fill",
						},
						"fillOne": {
							Summary: "An example of a query that fills IDs from s specific m2m field called cards",
							Value:   "$m2m=cards__id",
						},
					},
				},
				"q": {
					Name:          "$q",
					In:            "query",
					Description:   "Searches all string fields marked as Searchable",
					Required:      false,
					AllowReserved: true,
					Schema: &SchemaObject{
						Type:    "string",
						Default: "",
					},
				},
				"preload": {
					Name:          "$preload",
					In:            "query",
					Description:   "Fills the data from foreign keys",
					Required:      false,
					AllowReserved: true,
					Schema: &SchemaObject{
						Type:    "string",
						Default: "",
						Enum:    []interface{}{"", "true", "false"},
					},
					Examples: map[string]Example{
						"getDeleted": {
							Summary: "An example of a query that fills foreign key object",
							Value:   "$preload=true",
						},
					},
				},
				"next": {
					Name:          "$next",
					In:            "query",
					Description:   "Used in operation `method` to redirect the user to the specified path after the request. Value of `$back` will return the user back to the page",
					Required:      false,
					AllowReserved: true,
					Schema: &SchemaObject{
						Type:    "string",
						Default: "",
					},
				},
				"stat": {
					Name:          "$stat",
					In:            "query",
					Description:   "Returns the API call execution time in milliseconds",
					Required:      false,
					AllowReserved: true,
					Schema: &SchemaObject{
						Type:    "string",
						Default: "",
						Enum:    []interface{}{"", "true", "false"},
					},
					Examples: map[string]Example{
						"getStats": {
							Summary: "An example of a query that measures the execution time",
							Value:   "$stat=1",
						},
					},
				},
				"or": {
					Name:          "$or",
					In:            "query",
					Description:   "OR operator with multiple queries in the format of field=value. This `|` is used to separate the query parts and `+` is used for nested `AND` inside the the `OR` statement.",
					Required:      false,
					AllowReserved: true,
					Schema: &SchemaObject{
						Type:    "string",
						Default: "",
					},
					Examples: map[string]Example{
						"simple": {
							Summary: "An example of a query that returns records with active=1 or admin=1",
							Value:   "$or=active=1|admin=1",
						},
						"multiValueOr": {
							Summary: "An example of a query that returns records where the name starts with the letter a or ends with the letter a",
							Value:   "$or=name__startswith=a|name__endswith=a",
						},
						"nestedAnd": {
							Summary: "An example of a query that returns records with admin=1 or (active=1 and username=john)",
							Value:   "$or=admin=1|active=1+username=john",
						},
					},
				},
			},
			Responses: map[string]Response{
				"401": {
					Description: "Access denied due to missing authentication or insufficient permissions",
					Content: map[string]MediaType{
						"application/json": {
							Schema: &SchemaObject{
								Ref: "#/components/schemas/GeneralError",
							},
						},
					},
				},
			},
			SecuritySchemes: map[string]SecurityScheme{
				"apiKeyCookie": {
					Type: "apiKey",
					Name: "session",
					In:   "cookie",
				},
				"apiKeyQuery": {
					Type: "apiKey",
					Name: "session",
					In:   "query",
				},
				"JWT": {
					Type:         "http",
					Scheme:       "bearer",
					BearerFormat: "JWT",
				},
			},
		},
		Paths: getAuthPaths(),
		Security: []SecurityRequirement{
			{
				"apiKeyCookie": []string{},
				"apiKeyQuery":  []string{},
				"JWT":          []string{},
			},
		},
	}

	return s
}
