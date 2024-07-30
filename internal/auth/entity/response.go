package entity

import "alexdenkk/auth-api/model"

type LoginResponse struct {
	User  model.User `json:"user"`
	Token string     `json:"token"`
	Err   error      `json:"error"`
}
