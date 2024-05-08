package repository

import (
	"github.com/jmoiron/sqlx"
)

type TempPostgres struct {
	db *sqlx.DB
}

func NewTempPostgres(db *sqlx.DB) *TempPostgres {
	return &TempPostgres{db: db}
}

func (r *TempPostgres) Temp() {

}
