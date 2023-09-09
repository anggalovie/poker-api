package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	controller "api/src/controllers"
	middleware "api/src/middlewares"
)

func RegisterRoutes(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	UserprofileController := controller.UserprofileController{}

	router.HandleFunc("/update/profile/{id}", middleware.CheckAuth(UserprofileController.UpdateUserProfile(db))).Methods(http.MethodPut)
	router.HandleFunc("/profile/all", middleware.CheckAuth(UserprofileController.GetAllUserProfile(db))).Methods(http.MethodGet)
	router.HandleFunc("/profile", middleware.CheckAuth(UserprofileController.AddNewUserProfile(db))).Methods(http.MethodPost)
	router.HandleFunc("/profile/{id}", middleware.CheckAuth(UserprofileController.GetSingleUserProfile(db))).Methods(http.MethodGet)
	router.HandleFunc("/profile/{id}", middleware.CheckAuth(UserprofileController.DeleteSingleUserProfile(db))).Methods(http.MethodDelete)
	router.HandleFunc("/getprofile", middleware.CheckAuth(UserprofileController.GetProfile(db))).Methods(http.MethodGet)


	UserController := controller.UserController{}
	
	router.HandleFunc("/auth/login", UserController.LoginUser(db)).Methods(http.MethodPost)
	router.HandleFunc("/auth/signup", UserController.SignupUser(db)).Methods(http.MethodPost)

	return router
}
