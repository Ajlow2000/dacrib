package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/Ajlow2000/dacrib/handlers"

	"github.com/a-h/templ"
)


func main() {
    port := "3333"

	http.Handle("/", templ.Handler(handlers.Home()))
	// http.Handle("/articles", nil)
	// http.Handle("/experience", nil)
	// http.Handle("/about", nil)

    fmt.Printf("Initialized server\n")
    err := http.ListenAndServe(":" + port, nil)
    if errors.Is(err, http.ErrServerClosed) {
        fmt.Printf("server closed\n")
    } else if err != nil {
        fmt.Printf("error starting server: %s\n", err)
        os.Exit(1)
    }
}
