package handlers

import (
	"github.com/gin-gonic/gin"
	"haha/internal/models/DTO"
	"net/http"
)

func (h *Handler) getAllUser(c *gin.Context) {
	users, err := h.services.User.GetAllUsers()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, users)
}

func (h *Handler) getInfo(c *gin.Context) {
	id, _ := getUserId(c)

	user, err := h.services.User.GetUser(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) updateInfo(c *gin.Context) {
	id, _ := getUserId(c)

	var user DTO.UserUpdate

	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.User.UpdateUser(id, user); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
