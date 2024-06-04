package handlers

import (
	"net/http"

	"github.com/Ajlow2000/dacrib/components"
)

func Experience(w http.ResponseWriter, r *http.Request) {
    components.Experience().Render(r.Context(), w)
}
