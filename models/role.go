package models

import "main.go/app/role"

type MRole struct {
	ID   uint   `json:"id"`
	Role string `json:"role"`
}

func (mk *MRole) ToResponse() role.RoleResponseDTO {
	return role.RoleResponseDTO{
		ID: mk.ID,
		Role: mk.Role,
	}
}