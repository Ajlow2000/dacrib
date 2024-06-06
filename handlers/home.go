package handlers

import (
	"net/http"

	"github.com/Ajlow2000/dacrib/components"
)

func Home(w http.ResponseWriter, r *http.Request) {
	components.Home().Render(r.Context(), w)
}
