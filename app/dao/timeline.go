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
	timeline struct {
		db *sqlx.DB
	}
)

func NewTimeline(db *sqlx.DB) repository.Timeline {
	return &timeline{db: db}
}

func (r *timeline) GetTimeline(ctx context.Context, tlOptions *object.TimelineOptions) (*object.Timeline, error) {
	query := `
		SELECT
			s.id,
			s.content,
			s.create_at,
			a.id AS "account_id",
			a.username AS "account_username",
			a.create_at AS "account_create_at"
		FROM
			status s
		LEFT JOIN account a ON s.account_id = a.id
	`

	// `maxId` の値より小さいIDを持つstatusesを取得する
	if tlOptions.MaxID > 0 {
		// statusesのIDがmaxIdより小さいものを選択するクエリを追加
		query += fmt.Sprintf(" WHERE s.id < %d", tlOptions.MaxID)
	}

	// `sinceId` の値より大きいIDを持つstatusesを取得する
	if tlOptions.SinceID > 0 {
		if tlOptions.MaxID > 0 {
			query += " AND"
		} else {
			query += " WHERE"
		}
		query += fmt.Sprintf(" s.id > %d", tlOptions.SinceID)
	}

	query += fmt.Sprintf(" LIMIT %d", tlOptions.Limit)

	// 必要な情報をscanEntitiesにまとめる
	var scanEntities []struct {
		ID              int64     `db:"id"`
		Content         string    `db:"content"`
		CreateAt        time.Time `db:"create_at"`
		AccountID       int64     `db:"account_id"`
		Username        string    `db:"account_username"`
		AccountCreateAt time.Time `db:"account_create_at"`
	}

	err := r.db.SelectContext(ctx, &scanEntities, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find status from db: %w", err)
	}
	fmt.Println(scanEntities)

	var statuses []*object.Status
	for _, se := range scanEntities {
		statuses = append(statuses, &object.Status{
			ID:       se.ID,
			Content:  se.Content,
			CreateAt: se.CreateAt,
			Account: &object.Account{
				ID:       se.AccountID,
				Username: se.Username,
				CreateAt: se.AccountCreateAt,
			},
		})
	}

	// entityにStatusのスライスを格納する
	entity := &object.Timeline{
		Statuses: statuses,
	}

	return entity, nil

}
