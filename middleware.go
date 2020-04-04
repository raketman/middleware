package middleware

import (
	"github.com/gbrlsnchs/jwt/v3"
	"net/http"
)

type Middleware struct {
	TokenResolver TokenResolverContract
	ClientResolver ClientResolverContract
}

func (m Middleware) Handle(r *http.Request) Response {
	response := Response{Status: StatusError}

	token, err := m.TokenResolver.ResolveToken(r)

	if err != nil {
		response.Message = err.Error()
		return response
	}

	client, err := m.ClientResolver.ResolveClient(r)

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
