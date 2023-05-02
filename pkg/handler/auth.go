package handler

import (
	"alvile-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
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

func (h *Handler) createScheme(c *gin.Context) {
	var inputScheme models.Scheme

	if err := c.BindJSON(&inputScheme); err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	header := c.GetHeader("Authorization")
	email, errAuthorization := h.services.Registration.CheckAuthorization(header)
	if errAuthorization != nil {
		h.sendUnauthorized(c)
		return
	}

	output, err := h.services.Registration.CreateScheme(inputScheme, email)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}
	h.sendCreatedWithBody(c, output)

}

func (h *Handler) getScheme(c *gin.Context) {

	hed := c.GetHeader("Authorization")
	email, errAuthorization := h.services.Registration.CheckAuthorization(hed)
	if errAuthorization != nil {
		newErrorResponse(c, http.StatusUnauthorized, errAuthorization.Error())
		return
	}

	output, err := h.services.Registration.GetScheme(email)
	if err != nil {
		h.sendBadRequest(c, err.Error())
		h.sendInternalServerError(c)
		return
	}
	h.sendOKWithBody(c, output)
}
