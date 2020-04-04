Для обработки jwt нужно создать

middleware.Middleware

Для его создания необходимо 2 контракт:
 - TokenResolverContract
 - ClientResolverContract
 
 
Также в пакет входи контракт
 - AvailableClientResolverContract, который отвечает за получение списка клиентов

Пример реализаци, в пакете поставлются классы по умолчанию, которым реализует необходимые контракты

```

availableClientResolver = middleware.DefaultAvailableClientResolver{FilePath:"clients.json"}
tokenResolver = middleware.DefaultTokenResolver{}
clientResolver = middleware.DefaultClientResolver{
    AvailableClient: availableClientResolver,
    TokenResolver:   tokenResolver,
}

middlewareClient := middleware.Middleware{
    TokenResolver:  tokenResolver,
    ClientResolver: clientResolver,
}

``` 

