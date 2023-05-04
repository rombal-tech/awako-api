package models

type Scheme struct {
	Id           int64  `json:"id"`
	Name         string `json:"name" `
	Description  string `json:"description" `
	Author       string `json:"author" `
	CreationDate string `json:"creation_date"`
}

type SchemeOutput struct {
	List       []Scheme `json:"list"`
	TotalCount int64    `json:"total_count"`
}
