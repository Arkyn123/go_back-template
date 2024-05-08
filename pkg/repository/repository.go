package repository

import (
	"pet"

	"github.com/jmoiron/sqlx"
)

type Sueta interface {
	Create(sueta pet.Sueta) (int, int, error)
	FindAll() ([]pet.Sueta, error)
	Delete(id int) (pet.Sueta, error)
}

type Repository struct {
	Sueta
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Sueta: NewSuetaPostgres(db),
	}
}
