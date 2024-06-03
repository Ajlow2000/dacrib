package handlers

import (
	"net/http"

	"github.com/Ajlow2000/dacrib/components"
	"github.com/a-h/templ"
)

func Articles() http.HandlerFunc {
    return templ.Handler(components.Articles()).ServeHTTP
}

