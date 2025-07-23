package handler

import (
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
)

type Dependencies struct {
	AssetsFS http.FileSystem
}

type handlerFunc func(w http.ResponseWriter, r *http.Request) error

func RegisterRoutes(router *chi.Mux, deps Dependencies) {
	home := homeHandler{}

	router.Get("/", handler(home.handleIndex))

	router.Handle("/dist/*", http.StripPrefix("/dist/", http.FileServer(http.Dir("web/dist"))))
}

func handler(h handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			handleError(w, err)
		}
	}
}

func handleError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	slog.Error("error during request", slog.String("err", err.Error()))
}
