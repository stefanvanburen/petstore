package main

import (
	"net/http"

	"github.com/rs/cors"
	"github.com/stefanvanburen/petstore/gen/proto/go/pet/v1/petv1connect"
)

func main() {
	petServer := NewPetServer()
	mux := http.NewServeMux()
	path, handler := petv1connect.NewPetStoreServiceHandler(petServer)
	mux.Handle(path, handler)
	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	corsHandler := cors.Default().Handler(mux)
	http.ListenAndServe(":8080", corsHandler)
}
