package models

type Account struct {
	Deleted  bool   `json:"deleted"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password"binding:"required" `
}
