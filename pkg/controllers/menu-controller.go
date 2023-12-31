package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"golang.org/x/exp/slices"

	"strings"

	"github.com/gorilla/mux"
	"github.com/harin-h/nutritioncalculator/pkg/models"
	"github.com/harin-h/nutritioncalculator/pkg/utils"
)

var NewMenu models.RawMenu

func GetMenu(w http.ResponseWriter, r *http.Request) {
	newMenues := models.GetAllMenues()
	res, _ := json.Marshal(newMenues)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateMenu(w http.ResponseWriter, r *http.Request) {
	CreateMenu := &models.RawMenu{}
	utils.ParseBody(r, CreateMenu)
	fmt.Println(CreateMenu)
	fmt.Println(models.GetMenuPrimaryKey())
	CreateMenu.Id = models.GetMenuPrimaryKey() + 1
	fmt.Println(CreateMenu)
	CreateMenu.CreateMenu()
	w.WriteHeader(http.StatusOK)
}

func DeleteMenu(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	menuId := vars["menuId"]
	Id, err := strconv.ParseInt(menuId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	models.DeleteMenu(Id)
	w.WriteHeader(http.StatusOK)
}

func LikeMenu(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	menuId := vars["menuId"]
	Id, err := strconv.ParseInt(menuId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	userId := vars["userId"]
	menuDetails, db := models.GetMenuById(Id)
	userDetails, _ := models.GetUserById(userId)
	temp_slice := strings.Split(userDetails.FavoriteMenu, ",")
	if slices.Contains(temp_slice, menuId) {
		menuDetails.Like = menuDetails.Like - 1
	} else {
		menuDetails.Like = menuDetails.Like + 1
	}
	db.MustExec(`UPDATE menu SET count_like=$1 WHERE id=$2`, menuDetails.Like, Id)
	res, _ := json.Marshal(menuDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
