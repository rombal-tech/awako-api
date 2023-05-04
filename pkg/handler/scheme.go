package handler

import (
	"alvile-api/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createScheme(c *gin.Context) {
	var inputScheme models.Scheme

	if err := c.BindJSON(&inputScheme); err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	email, err := h.getAccountContext(c)
	if err != nil {
		h.sendInternalServerError(c)
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

	email, err := h.getAccountContext(c)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	output, err := h.services.Scheme.GetScheme(email)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}
	h.sendOKWithBody(c, output)
}
