package models

type RegistrationAccountInput struct {
	FirstName string `json:"firstName" binding:"required,excludesall=' ',printascii"`
	LastName  string `json:"lastName" binding:"required,excludesall=' ',printascii"`
	Email     string `json:"email" binding:"required,email,excludesall=' ',printascii"`
	Password  string `json:"password" binding:"required,excludesall=' ',printascii"`
}

type RegistrationAccountOutput struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
