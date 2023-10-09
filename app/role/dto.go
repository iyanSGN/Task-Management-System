package role

type RoleRequestDTO struct {
	ID 		uint		`json:"id"`
	Role	string		`json:"role" validate:"required"`
}

type RoleResponseDTO struct {
	ID		uint		`json:"id"`
	Role	string		`json:"role"`
}