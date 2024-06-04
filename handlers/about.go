package handlers

import (
	"net/http"

	"github.com/Ajlow2000/dacrib/components"
)


func About(w http.ResponseWriter, r *http.Request) {
    components.About().Render(r.Context(), w)
}

