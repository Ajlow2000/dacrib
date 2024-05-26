package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/Ajlow2000/dacrib/handlers"
)

func main() {

    port := "3333"

    http.HandleFunc("/", handlers.Home)
    http.HandleFunc("/about", handlers.About)
    http.HandleFunc("/articles", handlers.Articles)
    http.HandleFunc("/experience", handlers.Experience)

    fmt.Printf("Initialized server\n")
    err := http.ListenAndServe(":" + port, nil)
    if errors.Is(err, http.ErrServerClosed) {
        fmt.Printf("server closed\n")
    } else if err != nil {
        fmt.Printf("error starting server: %s\n", err)
        os.Exit(1)
    }
}
