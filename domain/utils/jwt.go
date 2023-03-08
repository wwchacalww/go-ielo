package utils

import (
	"time"

	"github.com/go-chi/jwtauth/v5"
)

func CreateJWToken(name, email, roles, avartar_url, secret string) (string, error) {
	jwtoken := jwtauth.New("HS256", []byte(secret), nil)

	exp := time.Now().Add(1 * time.Hour).Unix() // expiration time
	_, tokenString, err := jwtoken.Encode(map[string]interface{}{"name": name, "email": email, "role": roles, "avatar_url": avartar_url, "exp": exp})
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
