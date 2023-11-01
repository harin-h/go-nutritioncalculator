package models

import (
	"time"

	"github.com/harin-h/nutritioncalculator/pkg/config"
	"github.com/jmoiron/sqlx"
)

// var db *sqlx.DB / declared in menu.go

type User struct {
	UserId       string    `db:"user_id"`
	Password     string    `db:"password"`
	UserName     string    `db:"username"`
	Weight       float64   `db:"weight"`
	Protein      float64   `db:"protein"`
	Fat          float64   `db:"fat"`
	Carb         float64   `db:"carb"`
	FavoriteMenu string    `db:"favorite_menu"`
	CreateDate   time.Time `db:"created_date"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	var schema = `
	CREATE TABLE IF NOT EXISTS user_2023 (
		user_id VARCHAR,
		password VARCHAR,
		username VARCHAR,
		weight DECIMAL(10,2),
		protein DECIMAL(10,2),
		fat DECIMAL(10,2),
		carb DECIMAL(10,2),
		favorite_menu VARCHAR,
		created_date TIMESTAMP,
		PRIMARY KEY (user_id)
	)`
	db.MustExec(schema)
}

func (user *User) CreateUser() {
	const query = "INSERT INTO user_2023 (user_id,password,username,weight,protein,fat,carb,favorite_menu,created_date) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)"
	tx := db.MustBegin()
	tx.MustExec(query, user.UserId, user.Password, user.UserName, user.Weight, user.Protein, user.Fat, user.Carb, "", time.Now().UTC())
	tx.Commit()
}

func GetUserById(Id string) (User, *sqlx.DB) {
	var user User
	db.Get(&user, "SELECT * FROM user_2023 WHERE user_id=$1", Id)
	return user, db
}

func GetUserByUserName(UserName string) (User, *sqlx.DB) {
	var user User
	db.Get(&user, "SELECT * FROM user_2023 WHERE username=$1", UserName)
	return user, db
}
