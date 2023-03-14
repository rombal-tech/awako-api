package handler

import (
	"alvile-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) registration(c *gin.Context) {
	var input models.Account
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	email, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"email": email,
	})

}

func (h *Handler) authorization(c *gin.Context) {

}
