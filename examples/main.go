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
	tokenResolver = middleware.DefaultTokenResolver{}
	clientResolver = middleware.DefaultClientResolver{
		AvailableClient: availableClientResolver,
		TokenResolver:   tokenResolver,
	}

	middlewareClient := middleware.Middleware{
		TokenResolver:  tokenResolver,
		ClientResolver: clientResolver,
	}


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := middlewareClient.Handle(r)

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