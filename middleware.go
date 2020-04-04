package middleware

import (
	"fmt"
	//"github.com/gbrlsnchs/jwt"
	"net/http"
)

// Конфигурация
type Middleware struct {
	TokenResolver TokenResolverContract
	ClientResolver ClientResolverContract
}

//var cacheJwtVerificators map[string]jwt.Ver // чтобы не создавать постоянно

func (m Middleware) Handle(r *http.Request) Response {
	response := Response{Status: StatusError}

	token, err := m.TokenResolver.ResolveToken(r)

	if err != nil {
		response.Message = err.Error()
		return response
	}

	//jwt.Sign()

	client, err := m.ClientResolver.ResolveClient(r)

	if err != nil {
		response.Message = err.Error()
		return response
	}

	fmt.Println("TOKEN:", token.Token, "CLIENT", client)
	// Проверим токен (зараенн


	response.Status = StatusSuccess

	return response
}
