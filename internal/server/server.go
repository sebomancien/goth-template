package server

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/sebomancien/goth-template/internal/server/handlers"
	"github.com/sebomancien/goth-template/internal/server/middlewares"
)

type Config struct {
	Host string
	Port int
}

func Run(config *Config, files fs.FS) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", middlewares.Middlewares(handlers.HomeHandler))
	mux.Handle("GET /static/", http.FileServer(http.FS(files)))

	log.Printf("Listening on port http://%s:%d", config.Host, config.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Port), mux)
}
