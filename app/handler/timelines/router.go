package timelines

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"

	"github.com/go-chi/chi/v5"
)

// Implementation of handler
type handler struct {
	tr repository.Timeline
}

// Create Handler for `/v1/timelines/public`
func NewRouter(tr repository.Timeline) http.Handler {
	r := chi.NewRouter()

	h := &handler{tr}
	r.Get("/public", h.GetTimeline)

	return r
}
