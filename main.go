package main

import (
	"embed"
	"fmt"
	"log"

	"github.com/sebomancien/goth-template/config"
	"github.com/sebomancien/goth-template/internal/database"
	"github.com/sebomancien/goth-template/internal/models"
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

	err = test(db)
	if err != nil {
		log.Fatal(err)
	}

	err = server.Run(&cfg.Server, staticFS)
	if err != nil {
		log.Fatal(err)
	}
}

func test(db database.Database) error {
	err := db.AddLog(models.LogLevelInfo, "This is a test message")
	if err != nil {
		return err
	}

	err = db.AddLog(models.LogLevelInfo, "This is a another test message")
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
