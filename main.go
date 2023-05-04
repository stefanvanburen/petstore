package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"

	"buf.build/gen/go/acme/petapis/bufbuild/connect-go/pet/v1/petv1connect"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/jub0bs/fcors"
	"github.com/stefanvanburen/petstore/internal/server"
)

//go:embed README.md
var readmeMarkdown []byte

func main() {
	path, handler := petv1connect.NewPetStoreServiceHandler(server.NewPetServer())
	mux := http.NewServeMux()
	mux.Handle(path, handler)
	mux.HandleFunc("/", func(responseWriter http.ResponseWriter, _ *http.Request) {
		responseWriter.Write(markdownToHTML(readmeMarkdown))
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
		log.Fatal(fmt.Errorf("setting up CORS: %s", err))
	}
	http.ListenAndServe(":8080", cors(mux))
}

func markdownToHTML(markdownContent []byte) []byte {
	return markdown.Render(
		parser.New().Parse(markdownContent),
		html.NewRenderer(html.RendererOptions{Flags: html.CommonFlags}),
	)
}
