package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
    port := "3333"


    http.HandleFunc("/", helloHandler)
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
