package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Ajlow2000/dacrib/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	PORT := 3333

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.CleanPath)
	r.Use(middleware.Recoverer)

	r.Get("/", handlers.Home)
	r.Get("/home", handlers.Home)
	r.Get("/articles", handlers.Articles)
	r.Get("/articles/{title}", handlers.ShowArticle)
	r.Get("/experience", handlers.Experience)
	r.Get("/about", handlers.About)

	r.NotFound(handlers.NotFound)

	fmt.Println("Server started")
	http.ListenAndServe(":"+strconv.Itoa(PORT), r)
}
