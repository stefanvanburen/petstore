package main

import (
	"net/http"

	"github.com/rs/cors"
	"github.com/stefanvanburen/petstore/internal/server"
	"go.buf.build/bufbuild/connect-go/acme/petapis/pet/v1/petv1connect"
)

func main() {
	petServer := server.NewPetServer()
	mux := http.NewServeMux()
	path, handler := petv1connect.NewPetStoreServiceHandler(petServer)
	mux.Handle(path, handler)
	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST).
	corsHandler := cors.Default().Handler(mux)
	http.ListenAndServe(":8080", corsHandler)
}
