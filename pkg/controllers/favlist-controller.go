package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/harin-h/nutritioncalculator/pkg/models"
	"github.com/harin-h/nutritioncalculator/pkg/utils"
)

var NewFavList models.FavList

func GetFavList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	FavLists := models.GetFavList(userId)
	res, _ := json.Marshal(FavLists)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateFavList(w http.ResponseWriter, r *http.Request) {
	CreateFavList := &models.FavList{}
	utils.ParseBody(r, CreateFavList)
	favlist := CreateFavList.CreateFavList()
	res, _ := json.Marshal(favlist)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteFavList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	favlistId := vars["favlistId"]
	Id, err := strconv.ParseInt(favlistId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	models.DeleteFavList(Id)
	w.WriteHeader(http.StatusOK)
}

func UpdateFavList(w http.ResponseWriter, r *http.Request) {
	var updateFavList = &models.FavList{}
	utils.ParseBody(r, updateFavList)
	vars := mux.Vars(r)
	favlistId := vars["favlistId"]
	Id, err := strconv.ParseInt(favlistId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	favlistDetails, db := models.GetFavListById(Id)
	if updateFavList.Name != "" {
		favlistDetails.Name = updateFavList.Name
	}
	db.Save(favlistDetails)
	res, _ := json.Marshal(favlistDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
