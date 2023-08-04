package statuses

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *handler) FindStatus(w http.ResponseWriter, r *http.Request) {
	IDStr := chi.URLParam(r, "ID")
	id, err := strconv.ParseInt(IDStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(id)
	status, err := h.sr.FindByID(r.Context(), id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if status == nil {
		http.Error(w, "ステータスが見つかりませんでした", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
