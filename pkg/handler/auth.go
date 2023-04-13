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

	header := c.GetHeader("Authorization")
	email, errAuthorization := h.services.Registration.CheckAuthorization(header)
	if errAuthorization != nil {
		h.sendUnauthorized(c)
		return
	}

	if err := c.BindJSON(&inputScheme); err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}
	id, err := h.services.Registration.CreateScheme(inputScheme, email)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{ //точка
		"id": id,
	})

}

func (h *Handler) getScheme(c *gin.Context) {
	var inputScheme models.Scheme

	hed := c.GetHeader("Authorization")
	email, errAuthorization := h.services.Registration.CheckAuthorization(hed)
	if errAuthorization != nil {
		newErrorResponse(c, http.StatusUnauthorized, errAuthorization.Error())
		return
	}

	if err := c.BindJSON(&inputScheme); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	name, err := h.services.GetScheme(email)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"email": email,
		"name":  name,
	})
}
