package models

import "main.go/app/user"

type MUser struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (mk *MUser) ToResponse() user.UserResponseDTO{
	return	user.UserResponseDTO{
		ID: mk.ID,
		Email: mk.Email,
		Password: mk.Password,
	}
}