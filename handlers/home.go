package handlers

import (
	"net/http"

	"github.com/Ajlow2000/dacrib/components"
	"github.com/a-h/templ"
)


func Home() http.Handler {
    return templ.Handler(components.Home())
}

