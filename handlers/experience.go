package handlers

import (
	"io"
	"net/http"
)


func Experience(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "In EXPERIENCE page")
}

