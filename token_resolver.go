package middleware

import (
	"net/http"
	"strings"
)

type DefaultTokenResolver struct {}

func (t DefaultTokenResolver) ResolveToken(r *http.Request) (Token, error) {
	keys, ok := r.URL.Query()["access_token"]

	var token string

	if !ok || len(keys[0]) < 1 {
		// Пробуем найти в заголовках
		token = r.Header.Get("Authorization")

		splitToken := strings.Split(token, "Bearer ")
		if len(splitToken) == 2 {
			token = splitToken[1]
		}

		if len(token) == 0 {
			return Token{}, &Error{Message: "Не удалось найти токен"}
		}
	} else {
		// Query()["key"] will return an array of items,
		// we only want the single item.
		token = string(keys[0])
	}

	return Token{Token: token}, nil
}