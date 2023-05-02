package handler

import (
	"alvile-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createScheme(c *gin.Context) {
	var inputScheme models.Scheme

	if err := c.BindJSON(&inputScheme); err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	// Todo Нам будет приходить с фронта автор схемы ?
	// Если да, то принимать почту не нужно
	header := c.GetHeader("Authorization")
	email, errAuthorization := h.services.Account.CheckAuthorization(header)

	// Todo Мы перехватываем ошибки еще в repository и там записываем в логи, по сути эти обработчики бесполезны?
	if errAuthorization != nil {
		h.sendUnauthorized(c)
		return
	}

	output, err := h.services.Scheme.CreateScheme(inputScheme, email)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	h.sendCreatedWithBody(c, output)

}

func (h *Handler) getScheme(c *gin.Context) {

	hed := c.GetHeader("Authorization")
	email, errAuthorization := h.services.Account.CheckAuthorization(hed)
	if errAuthorization != nil {
		newErrorResponse(c, http.StatusUnauthorized, errAuthorization.Error())
		return
	}

	output, err := h.services.Scheme.GetScheme(email)
	if err != nil {
		h.sendBadRequest(c, err.Error())
		h.sendInternalServerError(c)
		return
	}
	h.sendOKWithBody(c, output)
}
