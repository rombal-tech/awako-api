package models

type Scheme struct {
	Id           int64  `json:"id"`
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description" binding:"required"`
	Author       string `json:"author" `
	CreationDate string `json:"creation_date"`
}

type SchemeOutput struct {
	Id           int64  `db:"id" json:"id"`
	Name         string `db:"name" json:"name"`
	Description  string `db:"description" json:"description"`
	Author       string `db:"author" json:"author" `
	CreationDate string `db:"creation_date" json:"creation_date"`
}
