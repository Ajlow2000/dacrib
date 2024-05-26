package handlers

import (
	"net/http"

	"github.com/Ajlow2000/dacrib/components"
	"github.com/a-h/templ"
)

func Experience() http.Handler {
    return templ.Handler(components.Experience())
}
