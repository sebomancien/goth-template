package handlers

import (
	"net/http"

	"github.com/sebomancien/goth-template/internal/templ"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	body := templ.Hello()
	templ.Layout("Home", body).Render(r.Context(), w)
}
