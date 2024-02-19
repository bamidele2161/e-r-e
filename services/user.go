package services

import (
	"e_real_estate/db"
	"e_real_estate/models"
	"errors"
	"fmt"
)

type UserService struct {
	serverDb *db.Database
}

// creating new instance of UserService
func NewUserService(db *db.Database) *UserService { 
return &UserService{serverDb: db}
}


func (s UserService) CreateUser(userPayload models.UserPayload) (models.UserPayload, error) {
	//check db if user already exists

	row := s.serverDb.Db.QueryRow("SELECT user where user_email = $1", userPayload.Email)

	if row != nil {
		return models.UserPayload{}, errors.New(("User already exists"))
	}

	
	result, err := s.serverDb.Db.Exec(`Insert into user (user_first_name, user_last_name, user_email, user_password) values ($1, $2, $3, $4)`, userPayload.FirstName, userPayload.LastName, userPayload.Email, userPayload.Password)
	
		if err != nil { 
			return models.UserPayload{}, errors.New("An error occured while creating user")
		}
	fmt.Println("err", err)

	rowAffected, err:= result.RowsAffected()
	if err != nil { 
		return models.UserPayload{}, errors.New("An error occured while creating user")
	}
	fmt.Println("err", err)

	if rowAffected != 1 {
		return models.UserPayload{}, errors.New("Error occured")
	}


	createdUser := models.UserPayload{}
	err = s.serverDb.Db.QueryRow("SELECT * FROM user WHERE user_email = $1", userPayload.Email).
	Scan(&createdUser.FirstName, &createdUser.LastName, &createdUser.Email, &createdUser.Password)
		
	if err != nil {
		return models.UserPayload{}, errors.New("Error occured")
}

	return createdUser, nil
}

func (s UserService) CreateUser2(user models.UserPayload) error {
	return nil
}

func User(data models.UserPayload) (users models.UserPayload, err error) {
	return 
}
func Login() (users []models.LoginPayload, err error) {
	return 
}