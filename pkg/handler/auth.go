package handler

import (
	"alvile-api/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) registration(c *gin.Context) {
	var input models.RegistrationInput

	if err := c.BindJSON(&input); err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	account, err := h.services.Registration.CreateAccount(input.Email, input.Password)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	outputSession, err := h.services.Registration.CreateSession(account.Email)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	output := models.RegistrationOutput{
		Email:   account.Email,
		Session: outputSession.Session,
	}

	h.sendCreatedWithBody(c, &output)
}

func (h *Handler) authorization(c *gin.Context) {
	var input models.AuthorizationInput

	if err := c.BindJSON(&input); err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	isExist, err := h.services.Account.IsExistByEmail(input.Email)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if !isExist {
		h.sendBadRequest(c, "account not found")
		return
	}

	session, err := h.services.Registration.CreateSession(input.Email)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	output := models.AuthorizationOutput{
		Email:   input.Email,
		Session: session.Session,
	}

	h.sendOKWithBody(c, &output)
}
