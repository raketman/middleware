####Пакет для реализации гибкого go get -u allmiddleware

"github.com/raketman/middleware"

#### Основной класс для проверка JWT-токена

middleware.Middleware

Для его создания необходимо 2 контракта:
 - TokenResolverContract - класс для получения токена из запроса
 - ClientResolverContract - класс для определения клиента
 
 
Также в пакет входит контракт
 - AvailableClientResolverContract, который отвечает за получение списка клиентов

Пример реализаци, в пакете поставлются классы по умолчанию, которым реализует необходимые контракты

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