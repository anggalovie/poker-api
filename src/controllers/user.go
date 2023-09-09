package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"

	//"time"
	//"fmt"

	"gorm.io/gorm"

	"api/src/helpers"
	"api/src/models"
)

// Authorization Key
var authKey = []byte("somesecret")

// Encryption Key
var encKey = []byte("someothersecret")

var store = sessions.NewCookieStore(authKey, encKey)

type UserController struct{}

type ResponseOutput struct {
	User  models.User
	Token string
}

func (u UserController) SignupUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		User := models.User{}
		json.NewDecoder(r.Body).Decode(&User)

		if len(User.Name) < 3 {
			error.ApiError(w, http.StatusBadRequest, "Name should be at least 3 characters long!")
			return
		}

		if len(User.Phone) < 3 {
			error.ApiError(w, http.StatusBadRequest, "Phone should be at least 3 characters long!")
			return
		}

		if len(User.Password) < 3 {
			error.ApiError(w, http.StatusBadRequest, "Password should be at least 3 characters long!")
			return
		}

		if len(User.Role) < 3 {
			error.ApiError(w, http.StatusBadRequest, "Roles should be at least 3 characters")
			return
		}

		if result := db.Create(&User); result.Error != nil {
			error.ApiError(w, http.StatusInternalServerError, "Failed To Add new User in database! \n"+result.Error.Error())
			return
		}

		payload := helpers.Payload{
			Phone: User.Phone,
			Id:    User.ID,
		}

		const cookieName = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
		cookieValue := strconv.FormatUint(uint64(User.ID), 10)
		//fmt.Printf("Set cookie %s=%s\n", cookieName, cookieValue)

		cookie := http.Cookie{Name: cookieName, Value: cookieValue, Path: "/"}
		http.SetCookie(w, &cookie)

		token, err := helpers.GenerateJwtToken(payload)
		if err != nil {
			error.ApiError(w, http.StatusInternalServerError, "Failed To Generate New JWT Token!")
			return
		}

		helpers.RespondWithJSON(w, ResponseOutput{
			Token: token,
			User:  User,
		})
	}
}

func (u UserController) LoginUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		User := models.User{}

		type Credentials struct {
			Phone    string
			Password string
		}
		credentials := Credentials{}
		json.NewDecoder(r.Body).Decode(&credentials)

		if len(credentials.Phone) < 3 {
			error.ApiError(w, http.StatusBadRequest, "Invalid Phone!")
			return
		}

		if len(credentials.Password) < 3 {
			error.ApiError(w, http.StatusBadRequest, "Invalid Password!")
			return
		}

		if results := db.Where("phone = ? AND password = ?", credentials.Phone, credentials.Password).First(&User); results.Error != nil || results.RowsAffected < 1 {
			error.ApiError(w, http.StatusNotFound, "Invalid Phone, Please Signup!")
			return
		}

		if User.Password != credentials.Password {
			error.ApiError(w, http.StatusNotFound, "Password not match")
			return
		}

		payload := helpers.Payload{
			Phone: User.Phone,
			Id:    User.ID,
		}

		const cookieName = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
		cookieValue := strconv.FormatUint(uint64(User.ID), 10)
		//fmt.Printf("Set cookie %s=%s\n", cookieName, cookieValue)

		cookie := http.Cookie{Name: cookieName, Value: cookieValue, Path: "/"}
		http.SetCookie(w, &cookie)

		token, err := helpers.GenerateJwtToken(payload)
		if err != nil {
			error.ApiError(w, http.StatusInternalServerError, "Failed To Generate New JWT Token!")
			return
		}

		helpers.RespondWithJSON(w, ResponseOutput{
			Token: token,
			User:  User,
		})
	}
}
