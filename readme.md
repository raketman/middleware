####Пакет для реализации гибкого middleware

"github.com/raketman/middleware"

#### Основной класс для проверка JWT-токена

middleware.Middleware

```
Метод 
Handle(tokenResolver TokenResolverContract, clientResolver ClientResolverContract) Response

type Response struct {
	Status string // статус операции
	Payload string // полезная нагрузка в json
	Message string // текст статуса
}

```


Для его создания необходимо 2 контракта:
 - TokenResolverContract - класс для получения токена из запроса
 - ClientResolverContract - класс для определения клиента
 
 
Также в пакет входит контракт
 - AvailableClientResolverContract, который отвечает за получение списка клиентов

Пример реализаци, в пакете поставляются классы по умолчанию, которые реализуют необходимые контракты

```
availableClientResolver = middleware.DefaultAvailableClientResolver{FilePath:"clients.json"}
tokenResolver = middleware.DefaultTokenResolver{Request: r}
clientResolver = middleware.DefaultClientResolver{AvailableClient: availableClientResolver, Request: r}

response := middlewareClient.Handle(tokenResolver, clientResolver)

if response.Status == middleware.StatusSuccess {
    w.Write([]byte(response.Payload))
} else {
    w.WriteHeader(http.StatusForbidden)
    w.Write([]byte(response.Message))
}

``` 

#####Примеры реализации с net/http можно найти в папке example

#### Поддерживаемые методы
|         | SHA-256            | SHA-384            | SHA-512            |
|:-------:|:------------------:|:------------------:|:------------------:|
| HMAC    |         Да         |         Да         |         Да         |