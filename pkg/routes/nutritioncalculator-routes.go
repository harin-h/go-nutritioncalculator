package routes

import (
	"github.com/gorilla/mux"
	"github.com/harin-h/nutritioncalculator/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {

	// Menu
	router.HandleFunc("/menu/", controllers.GetMenu).Methods("GET")
	router.HandleFunc("/menu/", controllers.CreateMenu).Methods("POST")
	router.HandleFunc("/menu/{menuId}", controllers.DeleteMenu).Methods("DELETE")
	router.HandleFunc("/menu/like/{menuId}/{userId}", controllers.LikeMenu).Methods("PUT")

	// User
	router.HandleFunc("/user/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/user_id/{userId}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user/user_name/{username}", controllers.GetUserName).Methods("GET")
	router.HandleFunc("/user/login", controllers.LoginUser).Methods("POST")
	router.HandleFunc("/user/goal/{userId}", controllers.UpdateUserGoal).Methods("PUT")
	router.HandleFunc("/user/fav_menu/{userId}/{menuId}", controllers.UpdateUserFavMenu).Methods("PUT")

	// Fav List
	router.HandleFunc("/fav_list/{userId}", controllers.GetFavList).Methods("GET")
	router.HandleFunc("/fav_list/", controllers.CreateFavList).Methods("POST")
	router.HandleFunc("/fav_list/{favlistId}", controllers.DeleteFavList).Methods("DELETE")
	router.HandleFunc("/fav_list/{favlistId}", controllers.UpdateFavList).Methods("PUT")
}
