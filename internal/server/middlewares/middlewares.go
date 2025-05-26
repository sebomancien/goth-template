package middlewares

import (
	"log"
	"net/http"
	"time"

	"github.com/sebomancien/goth-template/internal/context"
)

func logging(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		start := time.Now()
		handler(w, r)
		duration := time.Since(start)
		log.Printf("%.3fms\n", 1000*duration.Seconds())
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
