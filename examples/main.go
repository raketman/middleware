package main

import (
	"github.com/raketman/middleware"
	"log"
	"net/http"
)

var availableClientResolver middleware.AvailableClientResolverContract
var tokenResolver middleware.TokenResolverContract
var clientResolver middleware.ClientResolverContract

func main() {
	availableClientResolver = middleware.DefaultAvailableClientResolver{FilePath:"clients.json"}
	middlewareClient := middleware.Middleware{}


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tokenResolver = middleware.DefaultTokenResolver{Request: r}
		clientResolver = middleware.DefaultClientResolver{AvailableClient: availableClientResolver, Request: r}

		response := middlewareClient.Handle(tokenResolver, clientResolver)

		log.Println("Response:", response)
		if response.Status == middleware.StatusSuccess {
			w.Write([]byte(response.Payload))
		} else {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(response.Message))
		}

	})

	http.ListenAndServe(":9012", nil)

}