package repository

import (
	"net/http"
	"pet"

	"github.com/jmoiron/sqlx"
)

type SuetaPostgres struct {
	db *sqlx.DB
}

func NewSuetaPostgres(db *sqlx.DB) *SuetaPostgres {
	return &SuetaPostgres{db: db}
}

func (r *SuetaPostgres) Create(sueta pet.Sueta) (int, int, error) {
	var id int

	query := "INSERT INTO sueta (title, body) values ($1, $2) RETURNING id"

	row := r.db.QueryRow(query, sueta.Title, sueta.Body)
	if err := row.Scan(&id); err != nil {
		return 0, http.StatusInternalServerError, err
	}

	return id, http.StatusOK, nil
}

func (r *SuetaPostgres) FindAll() ([]pet.Sueta, error) {
	var suetas []pet.Sueta

	query := "SELECT * FROM sueta"

	if err := r.db.Select(&suetas, query); err != nil {
		return nil, err
	}
	if suetas == nil {
		return []pet.Sueta{}, nil
	}
	return suetas, nil
}

func (r *SuetaPostgres) Delete(id int) (pet.Sueta, error) {

	query := "DELETE FROM sueta WHERE id = $1"

	row := r.db.QueryRow(query, id)

	var sueta pet.Sueta
	err := row.Scan(&sueta)
	if err != nil {
		return sueta, err
	}

	return sueta, nil
}
