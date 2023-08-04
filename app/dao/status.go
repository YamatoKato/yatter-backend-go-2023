package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	status struct {
		db *sqlx.DB
	}
)

// Create status repository
func NewStatus(db *sqlx.DB) repository.Status {
	return &status{db: db}
}

func (r *status) AddStatus(status *object.Status) error {
	tx, _ := r.db.Begin()
	fmt.Println(status.Content)
	if _, err := tx.Exec("insert into status (account_id, content) values (?, ?)", status.AccountID, status.Content); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert account: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}

func (r *status) FindByID(ctx context.Context, id int64) (*object.Status, error) {
	query := `
		SELECT
			s.id,
			s.content,
			s.create_at,
			a.id AS "account_id",
			a.username AS "username",
			a.create_at AS "account_create_at"
		FROM
			status s
		LEFT JOIN account a ON s.account_id = a.id
		WHERE
			s.id = ?
	`

	var scanEntity struct {
		ID              int64     `db:"id"`
		Content         string    `db:"content"`
		CreateAt        time.Time `db:"create_at"`
		AccountID       int64     `db:"account_id"`
		Username        string    `db:"username"`
		AccountCreateAt time.Time `db:"account_create_at"`
	}

	// entity := new(object.Status)
	err := r.db.QueryRowxContext(ctx, query, id).StructScan(&scanEntity)
	fmt.Println(scanEntity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find status from db: %w", err)
	}

	entity := &object.Status{
		ID:        scanEntity.ID,
		AccountID: scanEntity.AccountID,
		Content:   scanEntity.Content,
		CreateAt:  scanEntity.CreateAt,
	}

	account := &object.Account{
		ID:       scanEntity.AccountID,
		Username: scanEntity.Username,
		CreateAt: scanEntity.AccountCreateAt,
	}

	entity.Account = account

	return entity, nil
}
