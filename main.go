package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

    "github.com/Ajlow2000/dacrib/components"

	"github.com/a-h/templ"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}


func main() {
	http.HandleFunc("/", getRoot)
	http.Handle("/hello", templ.Handler(components.Hello("Alec")))

    err := http.ListenAndServe(":3333", nil)
    if errors.Is(err, http.ErrServerClosed) {
        fmt.Printf("server closed\n")
    } else if err != nil {
        fmt.Printf("error starting server: %s\n", err)
        os.Exit(1)
    }
}
