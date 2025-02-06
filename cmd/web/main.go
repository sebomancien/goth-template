package main

import (
	"log"
	"net/http"

	"github.com/sebomancien/goth-template/internal/templ"
)

func main() {
	fs := http.FileServer(http.Dir("internal/static"))

	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Listening on port http://localhost:3000")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	body := templ.Hello()
	templ.Layout("Home", body).Render(r.Context(), w)
}
