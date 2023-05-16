package models

type Scheme struct {
	Id           int64  `json:"id"`
	Name         string `json:"name" `
	Description  string `json:"description" `
	Author       string `json:"author" `
	CreationDate string `json:"creation_date" db:"creation_date"`
}

type SchemeOutput struct {
	List       []Scheme `json:"list"`
	TotalCount int64    `json:"total_count"`
}

type InputSchemaParameters struct {
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
	Search string `json:"search"`
}
