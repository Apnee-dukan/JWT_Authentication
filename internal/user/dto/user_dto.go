package dto

import "golang.org/x/text/number"

type RegisterRequest struct {
	Username string        `json:"username"`
	Email    string        `json:"email"`
	Password string        `json:"password"`
	Phone    number.Option `json:"phone"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
