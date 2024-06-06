package handlers

import (
	"net/http"

	"github.com/Ajlow2000/dacrib/components"
	"github.com/a-h/templ"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	components.NotFound(templ.SafeURL("/")).Render(r.Context(), w)
}

func NotFoundGoArticles(w http.ResponseWriter, r *http.Request) {
	components.NotFound(templ.SafeURL("/articles")).Render(r.Context(), w)
}
