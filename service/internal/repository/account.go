package repository

import (
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"
	"tratnik.net/service/internal/model"
)

type IAccount interface {
	GetByID(ctx context.Context, accountID int64) (*model.Account, error)
}

var _ IAccount = (*Account)(nil)

type Account struct {
	db *sql.DB
}

func NewAccount(db *sql.DB) *Account {
	return &Account{
		db: db,
	}
}

func (r *Account) GetByID(ctx context.Context, accountID int64) (*model.Account, error) {
	query := `
		SELECT
			"id",
			"name",
			"is_active"
		FROM "account" "a"
		WHERE "id" = $1
	`

	row := r.db.QueryRowContext(ctx, query, accountID)
	account := &model.Account{}

	err := row.Scan(
		&account.ID,
		&account.Name,
		&account.IsActive,
	)
	switch err {
	case nil:
	case sql.ErrNoRows:
		return nil, ErrNoResults
	default:
		logrus.WithError(err).WithField("id", accountID).Error("Unable to fetch account")
		return nil, ErrUnknown
	}

	return account, nil
}
