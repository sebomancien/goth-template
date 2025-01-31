package main

import (
	"context"
	"log"
	"net/http"

	"github.com/sebomancien/goth-template/internal/templ"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ.Hello().Render(context.Background(), w)
	})

	log.Println("Listening on port http://localhost:3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
