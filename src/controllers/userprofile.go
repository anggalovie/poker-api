package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	//"strings"
	//"encoding/base64"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"api/src/helpers"
	"api/src/models"
)

type UserprofileController struct{}

var error = helpers.CustomError{}

func (f UserprofileController) GetAllUserProfile(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		UserProfile := []models.User{}

		if results := db.Find(&UserProfile); results.Error != nil {
			error.ApiError(w, http.StatusInternalServerError, "Failed To Fetch User Profile from database!")
			return
		}

		db.Raw("select name,email,username,'********' as password,'********' as role from users").Scan(&UserProfile)

		type Response struct {
			Success bool          `json:"success"`
			Status  int           `json:"status"`
			Data    []models.User `json:"data"`
		}
		resp := Response{Success: true, Status: 200, Data: UserProfile}

		helpers.RespondWithJSON(w, resp)
	}
}

func (f UserprofileController) GetSingleUserProfile(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		UserProfile := models.User{}

		if results := db.Where("id = ?", params["id"]).First(&UserProfile); results.Error != nil || results.RowsAffected < 1 {
			error.ApiError(w, http.StatusNotFound, "Didn't Find User Profile with id = "+params["id"])
			return
		}

		db.Raw("select name,email,username,'********' as password,'********' as role  from users").Scan(&UserProfile)

		type Response struct {
			Success bool        `json:"success"`
			Status  int         `json:"status"`
			Data    models.User `json:"data"`
		}
		resp := Response{Success: true, Status: 200, Data: UserProfile}

		helpers.RespondWithJSON(w, resp)
	}
}

func (f UserprofileController) AddNewUserProfile(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		const cookieName = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
		css, err := r.Cookie(cookieName)
		if err != nil {
			error.ApiError(w, http.StatusInternalServerError, "Error - login first")
			return
		} else {
			UserProfile := models.User{}
			db.First(&UserProfile, css.Value)

			if UserProfile.Role != "administrator" {
				error.ApiError(w, http.StatusInternalServerError, "accses denied admin only")
				return
			}

			UserProfileAdd := models.User{}
			json.NewDecoder(r.Body).Decode(&UserProfileAdd)

			if len(UserProfileAdd.Name) < 3 {
				error.ApiError(w, http.StatusBadRequest, "Name should be at least 3 characters !")
				return
			}

			if result := db.Create(&UserProfileAdd); result.Error != nil {
				error.ApiError(w, http.StatusInternalServerError, "Failed To Add new User Profile in database!")
				return
			}

			type Response struct {
				Success bool        `json:"success"`
				Status  int         `json:"status"`
				Data    models.User `json:"data"`
			}
			resp := Response{Success: true, Status: 200, Data: UserProfileAdd}

			helpers.RespondWithJSON(w, resp)

		}

	}
}

func (f UserprofileController) DeleteSingleUserProfile(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		const cookieName = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
		css, err := r.Cookie(cookieName)
		if err != nil {
			error.ApiError(w, http.StatusInternalServerError, "Error - login first")
			return
		} else {
			params := mux.Vars(r)
			UserProfile := models.User{}
			db.First(&UserProfile, css.Value)

			if UserProfile.Role != "administrator" {
				error.ApiError(w, http.StatusInternalServerError, "accses denied admin only")
				return
			}

			UserProfileDelete := models.User{}

			if results := db.Where("id = ?", params["id"]).First(&UserProfileDelete); results.Error != nil || results.RowsAffected < 1 {
				error.ApiError(w, http.StatusNotFound, "Didn't Find User Profile with id = "+params["id"])
				return
			}

			if results := db.Delete(&UserProfileDelete); results.Error != nil || results.RowsAffected < 1 {
				error.ApiError(w, http.StatusInternalServerError, "Failed to Delete User Profile from the database!")
				return
			}

			type Response struct {
				Success  bool   `json:"success"`
				Status   int    `json:"status"`
				Messages string `json:"messages"`
			}
			resp := Response{Success: true, Status: 200, Messages: "success deleted data id " + params["id"]}

			helpers.RespondWithJSON(w, resp)

		}
	}
}

func (f UserprofileController) UpdateUserProfile(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		const cookieName = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
		css, err := r.Cookie(cookieName)
		if err != nil {
			error.ApiError(w, http.StatusInternalServerError, "Error - login first")
			return
		} else {
			params := mux.Vars(r)
			UserProfile := models.User{}
			db.First(&UserProfile, css.Value)

			if UserProfile.Role != "administrator" {
				error.ApiError(w, http.StatusInternalServerError, "accses denied admin only")
				return
			}

			UserProfileUpdate := models.User{}

			if results := db.Where("id = ?", params["id"]).First(&UserProfileUpdate); results.Error != nil || results.RowsAffected < 1 {
				error.ApiError(w, http.StatusNotFound, "Didn't Find User Profile with id = "+params["id"])
				return
			}
			requestBody, _ := ioutil.ReadAll(r.Body)
			var person models.User
			json.Unmarshal(requestBody, &person)

			if resultsUpdate := db.Raw("UPDATE users SET name = ?, phone = ?, birth = ?, gender = ?, role = ? WHERE id = ? ",
				person.Name, person.Phone, person.Birth, person.Gender, person.Role, params["id"]).Scan(&person); resultsUpdate.Error != nil {
				error.ApiError(w, http.StatusNotFound, "changed data is the same as existing data ... contact administrator")
				return
			}

			db.Raw("select name,phone,birth,gender,'********' as password from users").Scan(&person)

			type Response struct {
				Success bool        `json:"success"`
				Status  int         `json:"status"`
				Data    models.User `json:"data"`
			}
			resp := Response{Success: true, Status: 200, Data: person}

			helpers.RespondWithJSON(w, resp)

		}

	}
}

func (f UserprofileController) GetProfile(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		const cookieName = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
		css, err := r.Cookie(cookieName)
		if err != nil {
			error.ApiError(w, http.StatusInternalServerError, "Error - login first")
			return
		} else {
			UserProfile := models.User{}
			db.First(&UserProfile, css.Value)

			type Response struct {
				Success bool        `json:"success"`
				Status  int         `json:"status"`
				Data    models.User `json:"data"`
			}
			resp := Response{Success: true, Status: 200, Data: UserProfile}

			helpers.RespondWithJSON(w, resp)
		}

	}
}
