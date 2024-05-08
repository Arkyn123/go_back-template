package repository

import (
	"github.com/jmoiron/sqlx"
)

type Temp interface {
	Temp()
}

type Repository struct {
	Temp
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Temp: NewTempPostgres(db),
	}
}
