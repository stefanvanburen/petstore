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
	"github.com/jba/templatecheck"
	"github.com/jub0bs/fcors"
	"github.com/stefanvanburen/petstore/internal/petstoreservice"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"rsc.io/markdown"
)

const defaultPort = "54321"

var (
	//go:embed README.md
	readmeMarkdown string
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
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	wrapperTemplate, err := template.New("index.html").Parse(htmlTemplate)
	if err != nil {
		return fmt.Errorf("parsing template: %s", err)
	}
	checkedTemplate, err := templatecheck.NewChecked[template.HTML](wrapperTemplate)
	if err != nil {
		return fmt.Errorf("creating checked template: %s", err)
	}
	readmeHTML := template.HTML(markdown.ToHTML(markdown.Parse(readmeMarkdown)))
	mux := http.NewServeMux()
	path, handler := petv1connect.NewPetStoreServiceHandler(petstoreservice.New())
	mux.Handle(path, handler)
	reflector := grpcreflect.NewStaticReflector(petv1connect.PetStoreServiceName)
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

	slog.InfoContext(context.Background(), "starting PetStore server", "port", port)

	return http.ListenAndServe(
		":"+port,
		h2c.NewHandler(cors(mux), &http2.Server{}),
	)
}
