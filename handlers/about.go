package handlers

import (
	"net/http"

	"github.com/Ajlow2000/dacrib/components"
	"github.com/a-h/templ"
)


func About() http.Handler {
    return templ.Handler(components.About())
}

