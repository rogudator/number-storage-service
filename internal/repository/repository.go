package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	NumberStorage
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		NumberStorage: NewNumberStoragePostgreSQL(db),
	}
}