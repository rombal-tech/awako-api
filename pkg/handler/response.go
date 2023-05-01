package handler

import (
	"github.com/execaus/exloggo"
	"github.com/gin-gonic/gin"
)

type error struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	exloggo.Errorf(message)
	c.AbortWithStatusJSON(statusCode, error{message})
}
