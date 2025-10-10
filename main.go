package main

import (
	"cmp"
	"context"
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"log/slog"
	"net/http"
	"os"

	"buf.build/gen/go/acme/petapis/connectrpc/go/pet/v1/petv1connect"
	connectcors "connectrpc.com/cors"
	"connectrpc.com/grpcreflect"
	"github.com/jba/templatecheck"
	"github.com/jub0bs/cors"
	"github.com/stefanvanburen/petstore/internal/petstoreservice"
	"rsc.io/markdown"
)

const defaultPort = "54321"

var (
	//go:embed README.md
	readmeMarkdown string
	//go:embed wrapper.html.gotmpl
	htmlTemplate string
)

func main() {
	if err := run(context.Background(), os.Stdout); err != nil {
		slog.Error("run", "error", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, out io.Writer) error {
	port := cmp.Or(os.Getenv("PORT"), defaultPort)

	wrapperTemplate, err := template.New("").Parse(htmlTemplate)
	if err != nil {
		return fmt.Errorf("parsing template: %s", err)
	}
	checkedTemplate, err := templatecheck.NewChecked[template.HTML](wrapperTemplate)
	if err != nil {
		return fmt.Errorf("creating checked template: %s", err)
	}
	readmeHTML := template.HTML(markdown.ToHTML((&markdown.Parser{}).Parse(readmeMarkdown)))

	// Allow access from Buf Studio.
	corsMiddleware, err := cors.NewMiddleware(cors.Config{
		Origins:         []string{"https://buf.build"},
		Methods:         connectcors.AllowedMethods(),
		RequestHeaders:  connectcors.AllowedHeaders(),
		ResponseHeaders: connectcors.ExposedHeaders(),
	})
	if err != nil {
		return fmt.Errorf("setting up CORS: %s", err)
	}

	mux := http.NewServeMux()
	{
		petStoreServicePath, petStoreServiceHandler := petv1connect.NewPetStoreServiceHandler(petstoreservice.New())
		mux.Handle(petStoreServicePath, corsMiddleware.Wrap(petStoreServiceHandler))
	}
	reflector := grpcreflect.NewStaticReflector(petv1connect.PetStoreServiceName)
	{
		reflectorv1Path, reflectorv1Handler := grpcreflect.NewHandlerV1(reflector)
		mux.Handle(reflectorv1Path, corsMiddleware.Wrap(reflectorv1Handler))
	}
	{
		reflectorv1alphaPath, reflectorv1alphaHandler := grpcreflect.NewHandlerV1Alpha(reflector)
		mux.Handle(reflectorv1alphaPath, corsMiddleware.Wrap(reflectorv1alphaHandler))
	}

	logger := slog.New(slog.NewTextHandler(out, nil))

	mux.HandleFunc("GET /{$}", func(responseWriter http.ResponseWriter, request *http.Request) {
		if err := checkedTemplate.Execute(responseWriter, readmeHTML); err != nil {
			logger.ErrorContext(request.Context(), "checkedTemplate.Execute", "error", err)
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	logger.InfoContext(ctx, "starting PetStore server", "port", port)

	protocols := &http.Protocols{}
	protocols.SetHTTP1(true)
	// For gRPC clients, it's convenient to support HTTP/2 without TLS.
	protocols.SetUnencryptedHTTP2(true)
	s := &http.Server{
		Addr:      ":" + port,
		Handler:   mux,
		Protocols: protocols,
	}
	return s.ListenAndServe()
}
