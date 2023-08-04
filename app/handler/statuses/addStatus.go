package statuses

import (
	"encoding/json"
	"log"
	"net/http"

	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/auth"
)

type AddRequest struct {
	Status string `json:"status"`
}

func (h *handler) AddStatus(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()
	// ログイン中のアカウントを取得する
	account := auth.AccountOf(r)
	log.Println(account.ID) // ここが取得できてない

	var req AddRequest
	// AddRequestに対して、statusのリクエストボディをデコードする
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	status := new(object.Status)
	status.Content = req.Status

	if account == nil {
		http.Error(w, "ユーザーが見つかりませんでした", http.StatusNotFound)
		return
	}
	status.AccountID = account.ID

	if err := h.sr.AddStatus(status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
