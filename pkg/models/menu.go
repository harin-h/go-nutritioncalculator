package models

import (
	"github.com/harin-h/nutritioncalculator/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Menu struct {
	Id      int64
	Name    string
	Protein float64
	Fat     float64
	Carb    float64
	Creator string
	Like    int64
}

type RawMenu struct {
	gorm.Model
	Name      string
	Protein   float64
	Fat       float64
	Carb      float64
	CreatorId string
	Like      int64
}

func init() {
	config.Connect()
	db = config.GetDB()
	// db.AutoMigrate(&RawMenu{})
	db.AutoMigrate(RawMenu{})
}

func (menu *RawMenu) CreateMenu() *RawMenu {
	// db.NewRecord(menu)
	// db.Create(&menu)
	db.Create(menu)
	return menu
}

func GetAllMenues() []Menu {
	var Menues []Menu
	db.Model(&RawMenu{}).Select("raw_menus.id, raw_menus.name, raw_menus.protein, raw_menus.fat, raw_menus.carb, users.user_name as creator, raw_menus.like").Joins("left join users on raw_menus.creator_id = users.user_id").Scan(&Menues)
	return Menues
}

func DeleteMenu(Id int64) RawMenu {
	var menu RawMenu
	db.Where("Id=?", Id).Delete(menu)
	return menu
}

func GetMenuById(Id int64) (RawMenu, *gorm.DB) {
	var menu RawMenu
	db.Where("Id=?", Id).Find(&menu)
	return menu, db
}
