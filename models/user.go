package models

type UserPayload struct {
	Id int `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserResponseData struct {
	Id int `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
type CreatedUserResponse struct {
	Message string `json:"message"`
	Token string `json:"token"`
	Data UserResponseData `json:"data"`
	StatusCode int `json:"statusCode"`
}

type ErrorResponse struct {
	Error string `json:"error"`
	StatusCode int `json:"statusCode"`
}
type LoginPayload struct {
	email string
	password string
}
type LoginResponse struct {
	email string
	password string
}

type Stock struct {
	StockID int64  `json:"stockid"`
	Name    string `json:"name"`
	Price   int64  `json:"price"`
	Company string `json:"company"`
}