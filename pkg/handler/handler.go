package handler

import (
	"alvile-api/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewService(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	auth := router.Group("/auth")
	{
		auth.POST("/registration", h.registration)
		auth.POST("/authorization", h.authorization)
	}
	schem := router.Group("/schem")
	{
		schem.GET("/get", h.getScheme)
		schem.POST("/post", h.postScheme)
	}

	return router
}
