package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Timeline interface {
	GetTimeline(ctx context.Context, tlOptions *object.TimelineOptions) (*object.Timeline, error)
}
