package middleware

import (
	"github.com/gbrlsnchs/jwt/v3"
)

type Middleware struct {
}

func (m Middleware) Handle(tokenResolver TokenResolverContract, clientResolver ClientResolverContract) Response {
	response := Response{Status: StatusError}

	token, err := tokenResolver.ResolveToken()

	if err != nil {
		response.Message = err.Error()
		return response
	}

	client, err := clientResolver.ResolveClient()

	if err != nil {
		response.Message = err.Error()
		return response
	}

	hs, errAlg := сreateAlg(client)

	if errAlg != nil {
		response.Message = err.Error()
		return response
	}

	var pl jwt.Payload

	_, err = jwt.Verify([]byte(token.Token), hs, &pl)
	if err != nil {
		response.Message = err.Error()
		return response
	}
	payload, err := getPayload(token)

	response.Status = StatusSuccess
	response.Payload = string(payload)

	return response
}
