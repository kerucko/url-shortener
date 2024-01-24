package handlers

import (
	"go_url_shortener/pkg/redis"
	"log"
	"net/http"
	"time"
)

func PostNewUrlHandler(db *redis.Storage, timeout time.Duration) http.HandlerFunc {
	op := "PostNewUrlHandler"
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(op)
		w.Write([]byte("hello"))
	}
}
