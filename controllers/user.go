package controllers

import (
	"e_real_estate/models"
	"e_real_estate/services"
	"encoding/json"
	"fmt"
	"net/http"
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
	var userPayload models.UserPayload
	err := json.NewDecoder(r.Body).Decode(&userPayload) 
if err != nil { 
	http.Error(w, "Invalid request", http.StatusBadRequest)
	return
}
createdUser, err := c.UserService.CreateUser(userPayload)
if err != nil {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	return
}
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(createdUser)


}
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}