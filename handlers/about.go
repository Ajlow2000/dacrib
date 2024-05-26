package handlers

import (
	"io"
	"net/http"
)


func About(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "In ABOUT page")
}

