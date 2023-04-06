package dtos

type LoginDTO struct {
	Token string  `json:"token"`
	User  UserDto `json:"user"`
}
