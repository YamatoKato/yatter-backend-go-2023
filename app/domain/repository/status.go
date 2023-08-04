package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Status interface {
	AddStatus(status *object.Status) error
	FindByID(ctx context.Context, id int64) (*object.Status, error)
}
