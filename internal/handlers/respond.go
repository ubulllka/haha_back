package handlers

import (
	"github.com/gin-gonic/gin"
	"haha/internal/models/DTO"
	"net/http"
)

func (h *Handler) respond(c *gin.Context) {

	var input DTO.RespondModel
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userRole, err := getUserRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	if err := h.services.CreateRespond(userRole, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}
