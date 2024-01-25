package main

import (
	"fmt"
	"go_url_shortener/internal/config"
	"go_url_shortener/internal/handlers"
	"go_url_shortener/pkg/redis"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg := config.MustLoad()

	dbPath := fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)
	log.Println(dbPath)
	db, err := redis.New(dbPath, cfg.Timeout)
	if err != nil {
		log.Fatal(err)
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Post("/newurl", handlers.PostNewUrlHandler(db, cfg.Timeout))

	log.Fatal(http.ListenAndServe(":"+cfg.Server.Port, router))
}
