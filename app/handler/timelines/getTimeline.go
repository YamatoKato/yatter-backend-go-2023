package timelines

import (
	"encoding/json"
	"log"
	"net/http"

	"yatter-backend-go/app/domain/object"
)

const LIMIT_DEFAULT = 40

func (h *handler) GetTimeline(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	queries := r.URL.Query()

	tlOptions, err := object.ParseTimelineOptions(queries)
	if err != nil {
		log.Printf("failed to parse timeline options: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	statuses, err := h.tr.GetTimeline(ctx, tlOptions)
	if err != nil {
		log.Printf("err: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(statuses.Statuses); err != nil {
		log.Printf("err: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
