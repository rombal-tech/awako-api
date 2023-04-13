package models

type Scheme struct {
	Id          int64  `json:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Author      string `json:"author" `
}
