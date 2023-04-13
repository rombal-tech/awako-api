package models

type Session struct {
	Id            int64  `json:"id"`
	Email         string `json:"email" binding:"required"`
	SessionString string `json:"session_string" binding:"required"`
}
type SessionOutput struct {
	SessionString string `json:"session_string"`
}
