package middleware

import (
	"encoding/json"
	"net/http"
)

var clients []Client
var err error

type DefaultClientResolver struct {
	AvailableClient AvailableClientResolverContract
	Request *http.Request
}

type innerPayload struct {
	Iss string
}

func (t DefaultClientResolver) ResolveClient () (Client, error)  {
	if len(clients) == 0 {
		clients = t.AvailableClient.GetClients()
	}

	tokenResolver := DefaultTokenResolver{Request:t.Request}

	token, err := tokenResolver.ResolveToken()

	if err != nil {
		return Client{}, err
	}

	payload, err := getPayload(token)

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