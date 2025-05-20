package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sebomancien/goth-template/config"
	"github.com/sebomancien/goth-template/internal/db"
	"github.com/sebomancien/goth-template/internal/middleware"
	"github.com/sebomancien/goth-template/internal/model"
	"github.com/sebomancien/goth-template/internal/templ"
)

const (
	DEFAULT_HTTP_PORT = "3000"
)

//go:embed static/*
var staticFS embed.FS

//go:embed migration/*
var migrationFS embed.FS

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.Open(&cfg.Database, migrationFS)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = test(db)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", middleware.Middlewares(homeHandler))
	mux.Handle("/static/", http.FileServer(http.FS(staticFS)))

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = DEFAULT_HTTP_PORT
	}

	log.Printf("Listening on port http://localhost:%s", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
	if err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	body := templ.Hello()
	templ.Layout("Home", body).Render(r.Context(), w)
}

func test(db db.Database) error {
	tables, err := db.GetTables()
	if err != nil {
		return err
	}

	fmt.Println("Tables")
	for _, v := range tables {
		fmt.Printf(" - %s\n", v)
	}

	err = db.AddLog(model.LogLevelInfo, "This is a test message")
	if err != nil {
		return err
	}

	err = db.AddLog(model.LogLevelInfo, "This is a another test message")
	if err != nil {
		return err
	}

	logs, err := db.GetLogs()
	if err != nil {
		return err
	}

	fmt.Println("Logs")
	for _, v := range logs {
		fmt.Printf(" - %v %d %s\n", v.Timestamp.Format("2006/01/02 15:04:05"), v.Level, v.Message)
	}

	return nil
}
