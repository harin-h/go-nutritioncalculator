package config

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	db *sqlx.DB
)

func Connect() {

	d, err := sqlx.Connect("postgres", "postgres://postgresql_nutritioncalculator_user:seJLjTpDqFZBN3vOpetgD9yMuIoXRCq2@dpg-cl01qrrjdq6s73b5qu00-a.singapore-postgres.render.com/postgresql_nutritioncalculator")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *sqlx.DB {
	return db
}
