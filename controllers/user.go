package controllers

import (
	"e_real_estate/models"
	"e_real_estate/services"
	"e_real_estate/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

type UserController struct { 
	UserService *services.UserService
}
func Test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Hello World!"))
}
type MyHandler struct {}

func NewUserController (service *services.UserService) *UserController {
	return &UserController{UserService :service}
}


func (c UserController) CreateAccount(w http.ResponseWriter, r *http.Request) {
	// parse user body into userPayload
	w.Header().Set("Content-Type", "application/json")
	var userPayload models.UserPayload
	err := json.NewDecoder(r.Body).Decode(&userPayload) 

	if err != nil { 
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	//validate the payload
		emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

		if !emailRegex.MatchString(userPayload.Email) {
			utils.RespondWithError(w, http.StatusBadRequest, "Please provide a valid email address")
			return
		} else if !utils.Validator(w, userPayload.Password, "Password", 6) || 
		!utils.Validator(w, userPayload.FirstName, "First Name", 3) || 
		!utils.Validator(w, userPayload.LastName, "Last Name", 3) {
			return
		}

		// hash password
		hashedPassword, _ := utils.HashPassword(userPayload.Password, 6)

		userPayload.Password = string(hashedPassword)

		createdUser, err := c.UserService.CreateUser(userPayload)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		
		json.NewEncoder(w).Encode(createdUser)

}


func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}