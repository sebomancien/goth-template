package main

import (
	"embed"
	"log"

	"github.com/sebomancien/goth-template/config"
	"github.com/sebomancien/goth-template/internal/database"
	"github.com/sebomancien/goth-template/internal/server"
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

	db, err := database.Open(&cfg.Database, migrationFS)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = server.Run(&cfg.Server, staticFS)
	if err != nil {
		log.Fatal(err)
	}
}
