package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"encoding/base64"
)

var clients []Client
var err error

type DefaultClientResolver struct {
	AvailableClient AvailableClientResolverContract
	TokenResolver TokenResolverContract
}

type innerPayload struct {
	Iss string
}

func (t DefaultClientResolver) ResolveClient (r *http.Request) (Client, error)  {
	if len(clients) == 0 {
		clients = t.AvailableClient.GetClients()
	}

	token, err := t.TokenResolver.ResolveToken(r)

	if err != nil {
		return Client{}, err
	}

	// Расшифруем токен возьмем оттуда iss
	log.Print("TOKEN:", token)

	splitToken := strings.Split(token.Token, ".")
	if len(splitToken) == 3 {
		log.Print("TOKEN:", splitToken[1])

		encoding := base64.RawURLEncoding
		decoded := make([]byte, encoding.DecodedLen(len(splitToken[1])))
		if _, err := encoding.Decode(decoded, []byte(splitToken[1])); err != nil {
			return Client{}, err
		}

		issPayload := innerPayload{Iss:""}
		json.Unmarshal(decoded, &issPayload)

		if len(issPayload.Iss) > 0 {
			// найдем клиент, если совпадает код
			for _, client := range clients {
				if client.Code ==  issPayload.Iss{
					return client, nil
				}
			}
		}
	}

	// найдем первый по умолчанию
	for _, client := range clients {
		if client.Default {
			return client, nil
		}
	}

	return Client{}, &Error{Message: "Не найден клиент"}
}