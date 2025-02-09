package middleware

import (
	"log"
	"net/http"

	"github.com/sebomancien/goth-template/internal/context"
)

func logging(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		handler(w, r)
	}
}

func authenticating(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.SetTheme(r.Context(), "dark")
		handler(w, r.WithContext(ctx))
	}
}

func Middlewares(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return authenticating(logging(handler))
}
