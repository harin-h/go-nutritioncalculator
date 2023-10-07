package controllers

import (
	"encoding/json"
	_ "fmt"
	"net/http"
	"sort"
	"strings"

	"golang.org/x/exp/slices"

	"github.com/gorilla/mux"
	"github.com/harin-h/nutritioncalculator/pkg/models"
	"github.com/harin-h/nutritioncalculator/pkg/utils"
)

var NewUser models.User

func CreateUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	u := CreateUser.CreateUser()
	res, _ := json.Marshal(u)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	userId := vars["userId"]
	userDetails, _ := models.GetUserById(userId)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUserGoal(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var updateUser = &models.User{}
	utils.ParseBody(r, updateUser)
	vars := mux.Vars(r)
	userId := vars["userId"]
	userDetails, db := models.GetUserById(userId)
	if updateUser.UserName != "" {
		userDetails.UserName = updateUser.UserName
	}
	if updateUser.Weight != 0 {
		userDetails.Weight = updateUser.Weight
	}
	if updateUser.Protein != 0 {
		userDetails.Protein = updateUser.Protein
	}
	if updateUser.Fat != 0 {
		userDetails.Fat = updateUser.Fat
	}
	if updateUser.Carb != 0 {
		userDetails.Carb = updateUser.Carb
	}
	db.Save(&userDetails)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUserFavMenu(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	menuId := vars["menuId"]
	userId := vars["userId"]
	userDetails, db := models.GetUserById(userId)
	var temp_favlist []string
	if userDetails.FavoriteMenu != "" {
		temp_favlist = strings.Split(userDetails.FavoriteMenu, ",")
	}
	if slices.Contains(temp_favlist, menuId) {
		index := slices.Index(temp_favlist, menuId)
		if index == 0 {
			userDetails.FavoriteMenu = strings.Join(temp_favlist[index+1:], ",")
		} else if index == len(temp_favlist)-1 {
			userDetails.FavoriteMenu = strings.Join(temp_favlist[:index], ",")
		} else {
			userDetails.FavoriteMenu = strings.Join(temp_favlist[:index], ",") + "," + strings.Join(temp_favlist[index+1:], ",")
		}
	} else {
		temp_favlist = append(temp_favlist, menuId)
		sort.Strings(temp_favlist)
		userDetails.FavoriteMenu = strings.Join(temp_favlist, ",")
	}
	db.Save(&userDetails)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
