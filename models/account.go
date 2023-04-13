package models

type AccountInput struct {
	Deleted  bool   `json:"deleted"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AccountOutput struct {
	Email string `json:"email" `
}
