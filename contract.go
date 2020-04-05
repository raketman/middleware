package middleware

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
	ResolveToken() (Token, error)
}

/**
Интерфейс для алгоритма определения клиента, который будет обрабатывать запрос
*/
type ClientResolverContract interface {
	ResolveClient() (Client, error)
}
