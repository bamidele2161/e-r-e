package models

type UserPayload struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
	Password  string `db:"password"`
}

type LoginPayload struct {
	email string
	password string
}
type LoginResponse struct {
	email string
	password string
}