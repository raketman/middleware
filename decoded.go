package middleware

import (
	"encoding/base64"
	"strings"
)

func getPayload (token Token) ([]byte, error) {
	splitToken := strings.Split(token.Token, ".")
	if len(splitToken) != 3 {
		return []byte(""), &Error{Message:"Incorrect token"}
	}

	return decoded(splitToken[1])
}

func decoded (value string) ([]byte, error) {
	encoding := base64.RawURLEncoding
	decoded := make([]byte, encoding.DecodedLen(len(value)))
	if _, err := encoding.Decode(decoded, []byte(value)); err != nil {
		return []byte(""), err
	}

	return decoded, nil
}