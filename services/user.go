package services

import (
	"e_real_estate/db"
	"e_real_estate/models"
	"e_real_estate/utils"
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

func (s UserService) CreateUser(userPayload models.UserPayload) (models.CreatedUserResponse, error) {
	//check db if user already exists


	row := s.serverDb.Db.QueryRow(`SELECT * FROM users WHERE email=$1`, userPayload.Email)
	existingUser := models.UserPayload{}

	err := row.Scan(
		&existingUser.Id,
		&existingUser.FirstName, 
		&existingUser.LastName,
		&existingUser.Email,
		&existingUser.Password )


	if err == nil { 
		return models.CreatedUserResponse{}, errors.New("User with the given email alreay exists")
	}

	_, err = s.serverDb.Db.Exec(`Insert into users (first_name, last_name, email, password) values ($1, $2, $3, $4)`, userPayload.FirstName, userPayload.LastName, userPayload.Email, userPayload.Password)
	
		if err != nil { 
			return models.CreatedUserResponse{}, errors.New("An error occured while creating user")
		}

	createdUser := models.UserPayload{}
	affectedRow := s.serverDb.Db.QueryRow(`SELECT * FROM users WHERE email=$1`, userPayload.Email)
	
	err = affectedRow.Scan(&createdUser.Id, &createdUser.FirstName, &createdUser.LastName, &createdUser.Email, &createdUser.Password)

	if err != nil {
		return models.CreatedUserResponse{}, errors.New("Error occured")
}
	token, err := utils.CreateToken(createdUser.Email)
	if err != nil {
		fmt.Println(err)
		return models.CreatedUserResponse{}, errors.New("Error occured")
	}
	responseData := models.UserResponseData{
		Email: createdUser.Email,
		FirstName : createdUser.FirstName,
		LastName :createdUser.LastName,
		Id :createdUser.Id,
	}
	

	response := models.CreatedUserResponse{
		Message: "User created successfully",
		Token: token,
		Data: responseData,
		StatusCode: 200,
	}
	return response, nil
	

}

func (s UserService) CreateUser2(user models.UserPayload) error {
	return nil
}

func User(data models.UserPayload) (users models.UserPayload, err error) {
	return 
}
func (s UserService) Login(userPayload models.LoginPayload) (models.CreatedUserResponse, error){

	row := s.serverDb.Db.QueryRow(`SELECT * FROM users WHERE email = $1`, userPayload.Email)
	existingUser := models.UserPayload{}

	err := row.Scan(
		&existingUser.Id,
		&existingUser.FirstName, 
		&existingUser.LastName,
		&existingUser.Email,
		&existingUser.Password )

	if err != nil {
		return models.CreatedUserResponse{}, errors.New("User not found")
	}

	
	//compare password
	err = utils.ComparePassword(userPayload.Password, []byte(existingUser.Password))
	if err != nil { 
		return models.CreatedUserResponse{}, errors.New("Invalid credentials")
	}

	//generate token
	token, err := utils.CreateToken(existingUser.Email)
	if err != nil { 
		fmt.Println(err)
		return models.CreatedUserResponse{}, errors.New("Error occured")}

		responseData := models.UserResponseData{
			Email: existingUser.Email,
			FirstName : existingUser.FirstName,
			LastName :existingUser.LastName,
			Id :existingUser.Id,
		}
		
		//return user data to client
		response := models.CreatedUserResponse{
			Message: "Login successfully",
			Token: token,
			Data: responseData,
			StatusCode: 200,
		}
	return response, nil
}