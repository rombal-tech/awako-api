package handler

import (
	"alvile-api/errors"
	"github.com/execaus/exloggo"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CheckAuthorization(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		h.sendUnauthorized(c)
		return
	}
	email, err := h.services.Account.CheckAuthorization(header)
	if err != nil {
		h.sendUnauthorized(c)
		return
	}
	h.setAccountContext(c, email)

	c.Next()
}

func (h *Handler) setAccountContext(c *gin.Context, email string) {
	c.Set("AccountEmail", email)
}

func (h *Handler) getAccountContext(c *gin.Context) (string, error) {
	email := c.GetString("AccountEmail")
	if email == "" {
		exloggo.Error("account context not found")
		return "", errors.ServerError
	}
	return email, nil
}
