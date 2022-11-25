package openapi

func getAuthPaths() map[string]Path {
	return map[string]Path{
		// Login auth API
		"/api/d/auth/login": {
			Summary:     "Login",
			Description: "Login API",
			Post: &Operation{
				Tags: []string{"Auth"},
				Responses: map[string]Response{
					"200": {
						Description: "Successful login",
						Content: map[string]MediaType{
							"application/json": {
								Schema: &SchemaObject{
									Type: "object",
									Properties: map[string]*SchemaObject{
										"status":  {Type: "string"},
										"jwt":     {Type: "string"},
										"session": {Type: "string"},
										"user": {
											Type: "object",
											Properties: map[string]*SchemaObject{
												"username":   {Type: "string"},
												"admin":      {Type: "boolean"},
												"first_name": {Type: "string"},
												"last_name":  {Type: "string"},
												"group_name": {Type: "string"},
											},
										},
									},
								},
							},
						},
					},
					"202": {
						Description: "Username and password are correct but MFA is required and OTP was not provided",
						Content: map[string]MediaType{
							"application/json": {
								Schema: &SchemaObject{
									Type: "object",
									Properties: map[string]*SchemaObject{
										"status":  {Type: "string"},
										"err_msg": {Type: "string"},
										"session": {Type: "string"},
									},
								},
							},
						},
					},
					"401": {
						Description: "Invalid credentials",
						Content: map[string]MediaType{
							"application/json": {
								Schema: &SchemaObject{
									Type: "object",
									Properties: map[string]*SchemaObject{
										"status":  {Type: "string"},
										"err_msg": {Type: "string"},
									},
								},
							},
						},
					},
				},
			},
			Parameters: []Parameter{
				{
					Name:        "username",
					In:          "query",
					Description: "Required for username/password login and single step MFA. But not required during the second step of a two-step MFA authentication",
					Schema: &SchemaObject{
						Type: "string",
					},
				},
				{
					Name:        "password",
					In:          "query",
					Description: "Required for username/password login and single step MFA. But not required during the second step of a two-step MFA authentication",
					Schema: &SchemaObject{
						Type: "string",
					},
				},
				{
					Name:        "otp",
					In:          "query",
					Description: "Not required for username/password login. Required for the second step in a two-step MFA and required single-step for MFA",
					Schema: &SchemaObject{
						Type: "string",
					},
				},
				{
					Name:        "session",
					In:          "query",
					Description: "Only required during the second step of a two-step MFA authentication",
					Schema: &SchemaObject{
						Type: "string",
					},
				},
			},
		},

		// Logout auth API
		"/api/d/auth/logout": {
			Summary:     "Logout",
			Description: "Logout API",
			Post: &Operation{
				Tags: []string{"Auth"},
				Responses: map[string]Response{
					"200": {
						Description: "Successful logout",
						Content: map[string]MediaType{
							"application/json": {
								Schema: &SchemaObject{
									Type: "object",
									Properties: map[string]*SchemaObject{
										"status": {Type: "string"},
									},
								},
							},
						},
					},
					"401": {
						Description: "User not logged in",
						Content: map[string]MediaType{
							"application/json": {
								Schema: &SchemaObject{
									Type: "object",
									Properties: map[string]*SchemaObject{
										"status":  {Type: "string"},
										"err_msg": {Type: "string"},
									},
								},
							},
						},
					},
				},
			},
			Parameters: []Parameter{
				{
					Ref: "#/components/parameters/CSRF",
				},
			},
		},

		// Signup auth API
		"/api/d/auth/signup": {
			Summary:     "Signup",
			Description: "Signup API",
			Post: &Operation{
				Tags: []string{"Auth"},
				Responses: map[string]Response{
					"200": {
						Description: "Successful signup",
						Content: map[string]MediaType{
							"application/json": {
								Schema: &SchemaObject{
									Type: "object",
									Properties: map[string]*SchemaObject{
										"status":  {Type: "string"},
										"jwt":     {Type: "string"},
										"session": {Type: "string"},
										"user": {
											Type: "object",
											Properties: map[string]*SchemaObject{
												"username":   {Type: "string"},
												"admin":      {Type: "boolean"},
												"first_name": {Type: "string"},
												"last_name":  {Type: "string"},
												"group_name": {Type: "string"},
											},
										},
									},
								},
							},
						},
					},
					"400": {
						Description: "Invalid or missing signup data. More about the error in err_msg.",
						Content: map[string]MediaType{
							"application/json": {
								Schema: &SchemaObject{
									Type: "object",
									Properties: map[string]*SchemaObject{
										"status":  {Type: "string"},
										"err_msg": {Type: "string"},
									},
								},
							},
						},
					},
				},
			},
			Parameters: []Parameter{
				{
					Name:        "username",
					In:          "query",
					Required:    true,
					Description: "Username can be any string or an email",
					Schema: &SchemaObject{
						Type: "string",
					},
				},
				{
					Name:        "password",
					In:          "query",
					Required:    true,
					Description: "Password",
					Schema: &SchemaObject{
						Type: "string",
					},
				},
				{
					Name:        "first_name",
					In:          "query",
					Required:    true,
					Description: "First name",
					Schema: &SchemaObject{
						Type: "string",
					},
				},
				{
					Name:        "last_name",
					In:          "query",
					Required:    true,
					Description: "Last name",
					Schema: &SchemaObject{
						Type: "string",
					},
				},
				{
					Name:        "email",
					In:          "query",
					Required:    true,
					Description: "Email",
					Schema: &SchemaObject{
						Type: "string",
					},
				},
			},
		},

		// Reset password auth API
		"/api/d/auth/resetpassword": {
			Summary:     "Reset Password",
			Description: "Reset Password API",
			Post: &Operation{
				Tags: []string{"Auth"},
				Responses: map[string]Response{
					"200": {
						Description: "Successful password reset",
						Content: map[string]MediaType{
							"application/json": {
								Schema: &SchemaObject{
									Type: "object",
									Properties: map[string]*SchemaObject{
										"status": {Type: "string"},
									},
								},
							},
						},
					},
					"202": {
						Description: "Password reset email sent",
						Content: map[string]MediaType{
							"application/json": {
								Schema: &SchemaObject{
									Type: "object",
									Properties: map[string]*SchemaObject{
										"status": {Type: "string"},
									},
								},
							},
						},
					},
					"400": {
						Description: "Missing password rest data. More about the error in err_msg.",
						Content: map[string]MediaType{
							"application/json": {
								Schema: &SchemaObject{
									Type: "object",
									Properties: map[string]*SchemaObject{
										"status":  {Type: "string"},
										"err_msg": {Type: "string"},
									},
								},
							},
						},
					},
					"401": {
						Description: "Invalid or expired OTP",
						Content: map[string]MediaType{
							"application/json": {
								Schema: &SchemaObject{
									Type: "object",
									Properties: map[string]*SchemaObject{
										"status":  {Type: "string"},
										"err_msg": {Type: "string"},
									},
								},
							},
						},
					},
					"403": {
						Description: "User does not have an email",
						Content: map[string]MediaType{
							"application/json": {
								Schema: &SchemaObject{
									Type: "object",
									Properties: map[string]*SchemaObject{
										"status":  {Type: "string"},
										"err_msg": {Type: "string"},
									},
								},
							},
						},
					},
					"404": {
						Description: "username or email do not match any active user",
						Content: map[string]MediaType{
							"application/json": {
								Schema: &SchemaObject{
									Type: "object",
									Properties: map[string]*SchemaObject{
										"status":  {Type: "string"},
										"err_msg": {Type: "string"},
									},
								},
							},
						},
					},
				},
			},
			Parameters: []Parameter{
				{
					Name:        "uid",
					In:          "query",
					Description: "Email or uid is required",
					Schema: &SchemaObject{
						Type: "string",
					},
				},
				{
					Name:        "email",
					In:          "query",
					Description: "Email or uid is required",
					Schema: &SchemaObject{
						Type: "string",
					},
				},
				{
					Name:        "password",
					In:          "query",
					Description: "New password which is required in the second step with the OTP",
					Schema: &SchemaObject{
						Type: "string",
					},
				},
				{
					Name:        "otp",
					In:          "query",
					Description: "OTP is required in the second step with a new password",
					Schema: &SchemaObject{
						Type: "string",
					},
				},
			},
		},

		// Change password auth API
		"/api/d/auth/changepassword": {
			Summary:     "Change Password",
			Description: "Change Password API",
			Post: &Operation{
				Tags: []string{"Auth"},
				Responses: map[string]Response{
					"200": {
						Description: "Successful password reset",
						Content: map[string]MediaType{
							"application/json": {
								Schema: &SchemaObject{
									Type: "object",
									Properties: map[string]*SchemaObject{
										"status": {Type: "string"},
									},
								},
							},
						},
					},
					"400": {
						Description: "Missing new password",
						Content: map[string]MediaType{
							"application/json": {
								Schema: &SchemaObject{
									Type: "object",
									Properties: map[string]*SchemaObject{
										"status":  {Type: "string"},
										"err_msg": {Type: "string"},
									},
								},
							},
						},
					},
					"401": {
						Description: "current password is invalid",
						Content: map[string]MediaType{
							"application/json": {
								Schema: &SchemaObject{
									Type: "object",
									Properties: map[string]*SchemaObject{
										"status":  {Type: "string"},
										"err_msg": {Type: "string"},
									},
								},
							},
						},
					},
				},
			},
			Parameters: []Parameter{
				{
					Ref: "#/components/parameters/CSRF",
				},
				{
					Name:        "old_password",
					In:          "query",
					Required:    true,
					Description: "Current user password",
					Schema: &SchemaObject{
						Type: "string",
					},
				},
				{
					Name:        "new_password",
					In:          "query",
					Required:    true,
					Description: "New password",
					Schema: &SchemaObject{
						Type: "string",
					},
				},
			},
		},
	}
}
