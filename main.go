package main

import (
	"embed"
	"log"

	"github.com/sebomancien/goth-template/internal/server"
)

//go:embed static/*
var staticFS embed.FS

func main() {
	cfg := server.Config{
		Host: "localhost",
		Port: 3000,
	}

	err := server.Run(&cfg, staticFS)
	if err != nil {
		log.Fatal(err)
	}
}
