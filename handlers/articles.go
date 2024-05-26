package handlers

import (
	"io"
	"net/http"
)


func Articles(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "in ARTICLES page")
}

