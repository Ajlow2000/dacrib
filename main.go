package main

import (
	"net/http"

	"github.com/Ajlow2000/dacrib/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
    r := chi.NewRouter()

    r.Use(middleware.Logger)

    r.Get("/", handlers.Home())
    r.Get("/home", handlers.Home())
    r.Get("/articles", handlers.Articles())
    r.Get("/articles/{title}", handlers.GetArticle())
    r.Get("/experience", handlers.Experience())
    r.Get("/about", handlers.About())



    r.NotFound(handlers.NotFound())

    http.ListenAndServe(":3333", r)
}
