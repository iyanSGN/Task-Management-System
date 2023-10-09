package user

type UserRequestDTO struct {
	ID				uint		`json:"id"`
	Username		string		`json:"username" validate:"required"`
	Email 			string		`json:"email" validate:"required"`
	Password		string		`json:"password" validate:"required"`
}

type UserResponseDTO struct {
	ID 			uint			`json:"id"`
	Email		string			`json:"email"`
	Password	string			`json:"password"`
}