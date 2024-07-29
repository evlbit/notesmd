package types

import "time"

type RegisterUserPayload struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=128"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
	Updated  time.Time `json:"-"`
	Created  time.Time `json:"-"`
}
