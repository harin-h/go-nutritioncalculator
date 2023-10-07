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

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func GetFavList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	FavLists := models.GetFavList(userId)
	res, _ := json.Marshal(FavLists)
	w.Header().Set("Content-Type", "pkglication/json")
	enableCors(&w)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateFavList(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	CreateFavList := &models.FavList{}
	utils.ParseBody(r, CreateFavList)
	f := CreateFavList.CreateFavList()
	res, _ := json.Marshal(f)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteFavList(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	favlistId := vars["favlistId"]
	Id, err := strconv.ParseInt(favlistId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	favlist := models.DeleteFavList(Id)
	res, _ := json.Marshal(favlist)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateFavList(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
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
	if updateFavList.List != "" {
		favlistDetails.List = updateFavList.List
	}
	db.Save(&favlistDetails)
	res, _ := json.Marshal(favlistDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
