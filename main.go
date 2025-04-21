package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sebomancien/goth-template/internal/middleware"
	"github.com/sebomancien/goth-template/internal/templ"
)

const (
	DEFAULT_HTTP_PORT = "3000"
)

//go:embed static/*
var staticFS embed.FS

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", middleware.Middlewares(homeHandler))
	mux.Handle("/static/", http.FileServer(http.FS(staticFS)))

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = DEFAULT_HTTP_PORT
	}

	log.Printf("Listening on port http://localhost:%s", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
	if err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	body := templ.Hello()
	templ.Layout("Home", body).Render(r.Context(), w)
}
