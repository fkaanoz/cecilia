package core

import (
	"context"
	"github.com/jmoiron/sqlx"
	"time"
)

type UserCore struct {
	db *sqlx.DB
}

func NewUserCore(db *sqlx.DB) *UserCore {
	return &UserCore{db: db}
}

func (uc UserCore) ReadSessionID(userID string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	row := uc.db.QueryRowContext(ctx, "SELECT sid FROM users WHERE user_id=$1", userID)
	if row.Err() != nil {
		return "", row.Err()
	}

	var sid string
	err := row.Scan(&sid)
	if err != nil {
		return "", err
	}

	return sid, nil
}
