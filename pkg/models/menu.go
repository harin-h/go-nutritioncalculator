package models

import (
	"time"

	"github.com/harin-h/nutritioncalculator/pkg/config"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type Menu struct {
	Id      int64   `db:"id"`
	Name    string  `db:"name"`
	Protein float64 `db:"protein"`
	Fat     float64 `db:"fat"`
	Carb    float64 `db:"carb"`
	Creator string  `db:"username"`
	Like    int64   `db:"count_like"`
}

type RawMenu struct {
	Id         int64     `db:"id"`
	Name       string    `db:"name"`
	Protein    float64   `db:"protein"`
	Fat        float64   `db:"fat"`
	Carb       float64   `db:"carb"`
	CreatorId  string    `db:"creator_id"`
	Like       int64     `db:"count_like"`
	Status     int64     `db:"status"`
	CreateDate time.Time `db:"created_date"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	var schema = `
	CREATE TABLE IF NOT EXISTS menu (
		id INTEGER,
		name VARCHAR,
		protein DECIMAL(10,2),
		fat DECIMAL(10,2),
		carb DECIMAL(10,2),
		creator_id VARCHAR,
		count_like INTEGER,
		status INTEGER,
		created_date TIMESTAMP,
		PRIMARY KEY (id)
	)`
	db.MustExec(schema)
}

func (menu *RawMenu) CreateMenu() {
	const query = "INSERT INTO menu (id,name,protein,fat,carb,creator_id,count_like,status,created_date) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)"
	tx := db.MustBegin()
	tx.MustExec(query, menu.Id, menu.Name, menu.Protein, menu.Fat, menu.Carb, menu.CreatorId, 0, 1, time.Now().UTC())
	tx.Commit()
}

func GetMenuPrimaryKey() int64 {
	menu := RawMenu{}
	if err := db.Get(&menu, "SELECT * FROM menu WHERE id = (SELECT MAX(id) FROM menu)"); err != nil {
		return 0
	}
	return menu.Id
}

func GetAllMenues() []Menu {
	var Menues []Menu
	db.Select(&Menues,
		`SELECT 
		menu.id,
		menu.name,
		menu.protein,
		menu.fat,
		menu.carb,
		user_2023.username,
		menu.count_like
	FROM menu INNER JOIN user_2023
		ON menu.creator_id = user_2023.user_id
	WHERE menu.status=1`)
	return Menues
}

func DeleteMenu(Id int64) {
	const query = "UPDATE menu SET status=0 WHERE id=$1"
	tx := db.MustBegin()
	tx.MustExec(query, Id)
	tx.Commit()
}

func GetMenuById(Id int64) (RawMenu, *sqlx.DB) {
	var menu RawMenu
	db.Get(&menu,
		`SELECT 
		*
	FROM menu
	WHERE id=$1`,
		Id)
	return menu, db
}
