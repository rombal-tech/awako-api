package models

type AccountRegistrationOutput struct {
	Email         string `json:"email" `
	SessionString string `json:"session_string"`
}
