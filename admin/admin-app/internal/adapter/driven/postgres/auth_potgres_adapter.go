package postgres

import (
	"admin/internal/core/port/driven"
	"database/sql"
)

type authRepo struct {
	authRepo driven.AuthPortDriven
	db       *sql.DB
}

func NewAuthRepo(db *sql.DB) *authRepo {

	return &authRepo{
		db: db,
	}

}
