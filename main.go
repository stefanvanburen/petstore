package main

import (
	"net/http"

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
	http.ListenAndServe(":8080", mux)
}
