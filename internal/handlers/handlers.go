package handlers

import (
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

func PostNewUrlHandler(db *redis.Client, timeout time.Duration) http.HandlerFunc {
	op := "PostNewUrlHandler"
	return func(w http.ResponseWriter, r *http.Request) {}
}
