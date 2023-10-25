package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	"buf.build/gen/go/acme/petapis/connectrpc/go/pet/v1/petv1connect"
	"connectrpc.com/grpcreflect"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/jba/templatecheck"
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
		slog.Error("run", "error", err)
		os.Exit(1)
	}
}

func run() error {
	wrapperTemplate, err := template.New("index.html").Parse(htmlTemplate)
	if err != nil {
		return fmt.Errorf("parsing template: %s", err)
	}
	checkedTemplate, err := templatecheck.NewChecked[template.HTML](wrapperTemplate)
	if err != nil {
		return fmt.Errorf("creating checked template: %s", err)
	}
	readmeHTML := template.HTML(string(markdownToHTML(readmeMarkdown)))
	path, handler := petv1connect.NewPetStoreServiceHandler(petstoreservice.New())
	reflector := grpcreflect.NewStaticReflector(petv1connect.PetStoreServiceName)
	mux := http.NewServeMux()
	mux.Handle(path, handler)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		if err := checkedTemplate.Execute(responseWriter, readmeHTML); err != nil {
			slog.ErrorContext(request.Context(), "checkedTemplate.Execute", "error", err)
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
			return
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
