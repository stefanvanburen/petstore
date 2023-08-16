package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"buf.build/gen/go/acme/petapis/connectrpc/go/pet/v1/petv1connect"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/jub0bs/fcors"
	"github.com/stefanvanburen/petstore/internal/petstoreservice"
)

var (
	//go:embed README.md
	readmeMarkdown []byte
	//go:embed wrapper.html.tmpl
	htmlTemplate string
)

func main() {
	if err := run(); err != nil {
		log.Printf("error: %s", err)
		os.Exit(1)
	}
}

func run() error {
	wrapperTemplate, err := template.New("index.html").Parse(htmlTemplate)
	if err != nil {
		return fmt.Errorf("parsing template: %s", err)
	}
	path, handler := petv1connect.NewPetStoreServiceHandler(petstoreservice.New())
	mux := http.NewServeMux()
	mux.Handle(path, handler)
	mux.HandleFunc("/", func(responseWriter http.ResponseWriter, _ *http.Request) {
		if err := wrapperTemplate.Execute(
			responseWriter,
			template.HTML(string(markdownToHTML(readmeMarkdown))),
		); err != nil {
			log.Printf("responseWriter.Write: %s", err)
		}
	})
	cors, err := fcors.AllowAccess(
		fcors.FromOrigins("https://buf.build"),
		fcors.WithMethods(http.MethodPost),
		fcors.WithRequestHeaders(
			"connect-protocol-version",
			"content-type",
		),
	)
	if err != nil {
		return fmt.Errorf("setting up CORS: %s", err)
	}
	// > ListenAndServe always returns a non-nil error.
	// Ignore it.
	_ = http.ListenAndServe(":8080", cors(mux))
	return nil
}

func markdownToHTML(markdownContent []byte) []byte {
	return markdown.Render(
		parser.New().Parse(markdownContent),
		html.NewRenderer(html.RendererOptions{Flags: html.CommonFlags}),
	)
}
