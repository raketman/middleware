package middleware

import (
	"net/http"
	"strings"
)

type DefaultTokenResolver struct {
	Request *http.Request
}

func (t DefaultTokenResolver) ResolveToken() (Token, error) {
	keys, ok := t.Request.URL.Query()["access_token"]


	var token string

	if !ok || len(keys[0]) < 1 {
		// Пробуем найти в заголовках
		token = t.Request.Header.Get("Authorization")

		splitToken := strings.Split(token, "Bearer ")
		if len(splitToken) == 2 {
			token = splitToken[1]
		}

		if len(token) == 0 {
			return Token{}, &Error{Message: "Не удалось найти токен"}
		}
	} else {
		token = string(keys[0])
	}

	return Token{Token: token}, nil
}