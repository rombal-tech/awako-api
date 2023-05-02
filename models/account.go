package models

type Account struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"-" db:"password"`
	Deleted  string `json:"-" db:"deleted"`
}

type RegistrationInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,max=32"`
}

type RegistrationOutput struct {
	Email   string `json:"email"`
	Session string `json:"session"`
}

type AuthorizationInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,max=32"`
}

type AuthorizationOutput struct {
	Email   string `json:"email"`
	Session string `json:"session"`
}
