package uadmin

import (
	"encoding/base64"
	"math/big"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func dAPIOpenIDCertHandler(w http.ResponseWriter, r *http.Request) {
	buf, err := os.ReadFile(".jwt-rsa-public.pem")
	if err != nil {
		w.WriteHeader(404)
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "Unable to load public certificate",
		})
		return
	}
	cert, err := jwt.ParseRSAPublicKeyFromPEM(buf)
	if err != nil {
		w.WriteHeader(404)
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "Unable to parse public certificate",
		})
		return
	}
	obj := map[string][]map[string]string{
		"keys": {
			{
				"kid": "1",
				"use": "sig",
				"kty": "RSA",
				"alg": "RS256",
				"e":   base64.RawURLEncoding.EncodeToString(big.NewInt(int64(cert.E)).Bytes()),
				"n":   base64.RawURLEncoding.EncodeToString(cert.N.Bytes()),
			},
		},
	}

	ReturnJSON(w, r, obj)
}
