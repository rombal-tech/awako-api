package models

type Session struct {
	Session string `json:"session_string" db:"session"`
	Email   string `json:"email" db:"email"`
}
