package handler

import (
	"github.com/fallenezer/Unidash_Goshka/internal/template/home"
	"net/http"
)

type homeHandler struct{}

func (h homeHandler) handleIndex(w http.ResponseWriter, r *http.Request) error {
	return home.Index().Render(r.Context(), w)
}
