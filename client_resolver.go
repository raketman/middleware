package middleware

import (
	"encoding/json"
	"net/http"
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

	payload, err := GetPayload(token)

	if err != nil {
		return Client{}, err
	}

	issPayload := innerPayload{Iss:""}
	json.Unmarshal(payload, &issPayload)

	if len(issPayload.Iss) > 0 {
		// найдем клиент, если совпадает код
		for _, client := range clients {
			if client.Code ==  issPayload.Iss{
				return client, nil
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