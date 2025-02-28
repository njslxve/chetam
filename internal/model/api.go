package model

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterResponse struct {
	Token string `json:"token,omitempty"`
}

type Error struct {
	Errors string `json:"errors"`
}
