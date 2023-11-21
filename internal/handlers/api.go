package handlers

import (
	"net/http"

	"github.com/Sn0wye/go-learn/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	r.Get("/ping", func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		w.Write([]byte("pong"))
	},
	)

	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}