package main

import (
	"go_url_shortener/internal/config"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg := config.MustLoad()

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	log.Fatal(http.ListenAndServe(":"+cfg.Server.Port, router))
}
