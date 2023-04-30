package main

import (
	"log"
	"net/http"

	"buf.build/gen/go/acme/petapis/bufbuild/connect-go/pet/v1/petv1connect"
	"github.com/jub0bs/fcors"
	"github.com/stefanvanburen/petstore/internal/server"
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
			"content-type",
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	http.ListenAndServe(":8080", cors(mux))
}
