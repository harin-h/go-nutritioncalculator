package models

import (
	"github.com/harin-h/nutritioncalculator/pkg/config"
	"github.com/jinzhu/gorm"
)

// var db *gorm.DB / declared in menu.go

type User struct {
	gorm.Model
	UserId       string
	Password     string
	UserName     string
	Weight       float64
	Protein      float64
	Fat          float64
	Carb         float64
	FavoriteMenu string
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(User{})
}

func (user *User) CreateUser() *User {
	db.Create(user)
	return user
}

func GetUserById(Id string) (User, *gorm.DB) {
	var user User
	db.Where("user_id=?", Id).Find(&user)
	return user, db
}

func GetUserByUserName(UserName string) (User, *gorm.DB) {
	var user User
	db.Where("user_name=?", UserName).Find(&user)
	return user, db
}
