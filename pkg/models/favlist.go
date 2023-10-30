package models

import (
	"github.com/harin-h/nutritioncalculator/pkg/config"
	"github.com/jinzhu/gorm"
)

// var db *gorm.DB / declared in menu.go

type FavList struct {
	gorm.Model
	UserId  string
	Name    string
	Menues  string
	List    string
	Protein float64
	Fat     float64
	Carb    float64
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(FavList{})
}

func GetFavList(Id string) []FavList {
	var Favlists []FavList
	db.Where("user_id=?", Id).Find(&Favlists)
	return Favlists
}

func (favlist *FavList) CreateFavList() *FavList {
	db.Create(favlist)
	return favlist
}

func DeleteFavList(Id int64) {
	db.Where("ID=?", Id).Delete(FavList{})
}

func GetFavListById(Id int64) (FavList, *gorm.DB) {
	var favlist FavList
	db.Where("ID=?", Id).Find(&favlist)
	return favlist, db
}
