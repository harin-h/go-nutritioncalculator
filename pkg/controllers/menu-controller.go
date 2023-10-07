package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"golang.org/x/exp/slices"

	"github.com/gorilla/mux"
	"github.com/harin-h/nutritioncalculator/pkg/models"
	"github.com/harin-h/nutritioncalculator/pkg/utils"
)

var NewMenu models.RawMenu

func GetMenu(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	newMenues := models.GetAllMenues()
	res, _ := json.Marshal(newMenues)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateMenu(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	CreateMenu := &models.RawMenu{}
	utils.ParseBody(r, CreateMenu)
	m := CreateMenu.CreateMenu()
	res, _ := json.Marshal(m)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteMenu(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	menuId := vars["menuId"]
	Id, err := strconv.ParseInt(menuId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	menu := models.DeleteMenu(Id)
	res, _ := json.Marshal(menu)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateMenu(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var updateMenu = &models.RawMenu{}
	utils.ParseBody(r, updateMenu)
	vars := mux.Vars(r)
	menuId := vars["menuId"]
	Id, err := strconv.ParseInt(menuId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	menuDetails, db := models.GetMenuById(Id)
	if updateMenu.Name != "" {
		menuDetails.Name = updateMenu.Name
	}
	if updateMenu.Protein != 0 {
		menuDetails.Protein = updateMenu.Protein
	}
	if updateMenu.Fat != 0 {
		menuDetails.Fat = updateMenu.Fat
	}
	if updateMenu.Carb != 0 {
		menuDetails.Carb = updateMenu.Carb
	}
	menuDetails.Like = 0
	db.Save(&menuDetails)
	res, _ := json.Marshal(menuDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func LikeMenu(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	menuId := vars["menuId"]
	Id, err := strconv.ParseInt(menuId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	userId := vars["userId"]
	menuDetails, db := models.GetMenuById(Id)
	userDetails, _ := models.GetUserById(userId)
	temp_slice := []byte(userDetails.FavoriteMenu)
	if slices.Contains(temp_slice, []byte(menuId)[0]) {
		menuDetails.Like = menuDetails.Like - 1
	} else {
		menuDetails.Like = menuDetails.Like + 1
	}
	db.Save(&menuDetails)
	res, _ := json.Marshal(menuDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
