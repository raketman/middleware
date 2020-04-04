package middleware

import (
	"net/http"
)

/**
Интерфейс для получения списка доступных клиентов
 */
type AvailableClientResolverContract interface {
	GetClients() []Client
}

/**
Интерфейс для получения токена из запроса
*/
type TokenResolverContract interface {
	ResolveToken(r *http.Request) (Token, error)
}

/**
Интерфейс для алгоритма определения клиента, который будет обрабатывать запрос
*/
type ClientResolverContract interface {
	ResolveClient(r *http.Request) (Client, error)
}