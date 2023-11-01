package models

import (
	"time"

	"github.com/harin-h/nutritioncalculator/pkg/config"
	"github.com/jmoiron/sqlx"
)

// var db *splx.DB / declared in menu.go

type FavList struct {
	Id         int64     `db:"id"`
	UserId     string    `db:"user_id"`
	Name       string    `db:"name"`
	Menues     string    `db:"menues"`
	List       string    `db:"list"`
	Protein    float64   `db:"protein"`
	Fat        float64   `db:"fat"`
	Carb       float64   `db:"carb"`
	Status     int64     `db:"status"`
	CreateDate time.Time `db:"created_date"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	var schema = `
	CREATE TABLE IF NOT EXISTS fav_list (
		id INTEGER,
		user_id VARCHAR,
		name VARCHAR,
		menues VARCHAR,
		list VARCHAR,
		protein DECIMAL(10,2),
		fat DECIMAL(10,2),
		carb DECIMAL(10,2),
		Status INTEGER,
		created_date TIMESTAMP,
		PRIMARY KEY (id)
	)`
	db.MustExec(schema)
}

func (favlist *FavList) CreateFavList() {
	const query = "INSERT INTO fav_list (id,user_id,name,menues,list,protein,fat,carb,status,created_date) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)"
	tx := db.MustBegin()
	tx.MustExec(query, favlist.Id, favlist.UserId, favlist.Name, favlist.Menues, favlist.List, favlist.Protein, favlist.Fat, favlist.Carb, 1, time.Now().UTC())
	tx.Commit()
}

func GetFavoriteListPrimaryKey() int64 {
	favlist := FavList{}
	if err := db.Get(&favlist, "SELECT * FROM fav_list WHERE id = (SELECT MAX(id) FROM fav_list)"); err != nil {
		return 0
	}
	return favlist.Id
}

func GetFavList(Id string) []FavList {
	var Favlists []FavList
	db.Select(&Favlists, "SELECT * FROM fav_list WHERE user_id=$1 AND status=1 ORDER BY id", Id)
	return Favlists
}

func DeleteFavList(Id int64) {
	db.MustExec("UPDATE fav_list SET status=0 WHERE id=$1", Id)
}

func GetFavListById(Id int64) (FavList, *sqlx.DB) {
	var favlist FavList
	db.Get(&favlist, "SELECT * FROM fav_list WHERE id=$1", Id)
	return favlist, db
}
