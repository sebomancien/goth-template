package main

import (
	"embed"
	"log"

	"github.com/sebomancien/goth-template/config"
	"github.com/sebomancien/goth-template/internal/server"
)

//go:embed static/*
var staticFS embed.FS

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	err = server.Run(&cfg.Server, staticFS)
	if err != nil {
		log.Fatal(err)
	}
}
