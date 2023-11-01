package controllers

import (
	"encoding/json"
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
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	CreateUser.CreateUser()
	w.WriteHeader(http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	userDetails, _ := models.GetUserById(userId)
	userDetails.Password = ""
	res, _ := json.Marshal(userDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	userDetails, _ := models.GetUserByUserName(username)
	res, _ := json.Marshal(userDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var RequestUserLogin = &models.User{}
	utils.ParseBody(r, RequestUserLogin)
	userDetails, _ := models.GetUserById(RequestUserLogin.UserId)
	var responseData map[string]bool
	if RequestUserLogin.Password == userDetails.Password && len(RequestUserLogin.Password) > 0 {
		responseData = map[string]bool{
			"IsPasswordCorrect": true,
		}
	} else {
		responseData = map[string]bool{
			"IsPasswordCorrect": false,
		}
	}
	res, _ := json.Marshal(responseData)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUserGoal(w http.ResponseWriter, r *http.Request) {
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
	tx := db.MustBegin()
	tx.MustExec(`UPDATE user_2023 SET username=$1, weight=$2, protein=$3, fat=$4, carb=$5 WHERE user_id=$6`, userDetails.UserName, userDetails.Weight, userDetails.Protein, userDetails.Fat, userDetails.Carb, userDetails.UserId)
	tx.Commit()
	res, _ := json.Marshal(userDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUserFavMenu(w http.ResponseWriter, r *http.Request) {
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
	tx := db.MustBegin()
	tx.MustExec(`UPDATE user_2023 SET favorite_menu=$1 WHERE user_id=$2`, userDetails.FavoriteMenu, userDetails.UserId)
	tx.Commit()
	res, _ := json.Marshal(userDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
