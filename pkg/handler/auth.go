package handler

import (
	"alvile-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) registration(c *gin.Context) {
	var input models.AccountInput
	var session models.Session
	if err := c.BindJSON(&input); err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}
	output, err := h.services.Registration.CreateUser(input)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}
	outputSession, err := h.services.Registration.CreateSession(&session, input.Email, input.Password)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}
	output.SessionString = outputSession.SessionString

	if err != nil {
		h.sendInternalServerError(c)
		return
	}
	h.sendCreatedWithBody(c, output)

}

func (h *Handler) authorization(c *gin.Context) {
	var inputAccount models.AccountInput
	var session models.Session
	if err := c.BindJSON(&inputAccount); err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}
	output, err := h.services.Registration.CreateSession(&session, inputAccount.Email, inputAccount.Password)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}
	h.sendCreatedWithBody(c, output)

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
