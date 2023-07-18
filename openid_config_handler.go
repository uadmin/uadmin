package uadmin

import "net/http"

func JWTConfigHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"issuer":                 JWTIssuer,
		"authorization_endpoint": JWTIssuer + "/api/d/auth/openidlogin",
		"token_endpoint":         "",
		"userinfo_endpoint":      JWTIssuer + "/api/d/auth/userinfo",
		"jwks_uri":               JWTIssuer + "/api/d/auth/certs",
		"scopes_supported": []string{
			"openid",
			"email",
			"profile",
		},
		"response_types_supported": []string{
			"code",
			"token",
			"id_token",
			"code token",
			"code id_token",
			"token id_token",
			"code token id_token",
			"none",
		},
		"subject_types_supported": []string{
			"public",
		},
		"id_token_signing_alg_values_supported": []string{
			"RS256",
		},
		"claims_supported": []string{
			"aud",
			"email",
			"email_verified",
			"exp",
			"family_name",
			"given_name",
			"iat",
			"iss",
			"locale",
			"name",
			"picture",
			"sub",
		},
	}

	ReturnJSON(w, r, data)
}
