package main

import (
	"log"
	"net/http"

	"github.com/jub0bs/fcors"
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
	cors, err := fcors.AllowAccess(
		fcors.FromOrigins("https://studio.buf.build"),
		fcors.WithMethods(http.MethodPost),
		fcors.WithRequestHeaders(
			"connect-protocol-version",
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	http.ListenAndServe(":8080", cors(mux))
}
