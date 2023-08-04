package accounts

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Request body for `GET /v1/accounts/:username`
type FindAccountRequest struct {
	Username string
}

// Handle request for `GET /v1/accounts/:username`
func (h *handler) FindAccount(w http.ResponseWriter, r *http.Request) {
	// ユーザ名からユーザを取得
	account, err := h.ar.FindByUsername(r.Context(), chi.URLParam(r, "username"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if account == nil {
		http.Error(w, "ユーザーが見つかりませんでした", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
