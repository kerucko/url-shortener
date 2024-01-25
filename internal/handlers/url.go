package handlers

import (
	"go_url_shortener/internal/entities"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator"
)

type Storage interface {
	SaveUrl(saveUrl string, alias string, expiration time.Duration, timeout time.Duration) error
	GetUrl(alias string, timeout time.Duration) (string, error)
	DeleteUrl(alias string, timeout time.Duration) error
}

func PostNewUrlHandler(db Storage, timeout time.Duration) http.HandlerFunc {
	op := "PostNewUrlHandler"
	return func(w http.ResponseWriter, r *http.Request) {
		requestId := middleware.GetReqID(r.Context())
		log.Printf("op: %s; RequestID: %s\n", op, requestId)

		var request entities.Request
		err := render.DecodeJSON(r.Body, &request)
		if err != nil {
			log.Printf("%s; %s; cannot decode request json %s", op, requestId, err)
			render.JSON(w, r, entities.Error("cannot decode request json"))
			return
		}
		log.Printf("%s; %s; request JSON: %v\n", op, requestId, request)

		if err := validator.New().Struct(request); err != nil {
			log.Printf("%s, %s; invalid request %s", op, requestId, err)
			render.JSON(w, r, entities.Error("invalid request"))
			return
		}
	}
}
