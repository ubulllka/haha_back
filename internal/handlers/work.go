package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"haha/internal/models"
	"haha/internal/models/DTO"
	"net/http"
	"strconv"
	"strings"
)

func (h Handler) createWork(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	userRole, err := getUserRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	if !strings.EqualFold(userRole, models.APPLICANT) && !strings.EqualFold(userRole, models.ADMIN) {
		newErrorResponse(c, http.StatusForbidden, errors.New("not enough rights").Error())
		return
	}

	resumeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid resume id param")
		return
	}

	var work DTO.WorkCreate

	if err := c.BindJSON(&work); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Work.CreateWork(userId, uint(resumeId), userRole, work)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h Handler) updateWork(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	userRole, err := getUserRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	if !strings.EqualFold(userRole, models.APPLICANT) && !strings.EqualFold(userRole, models.ADMIN) {
		newErrorResponse(c, http.StatusForbidden, errors.New("not enough rights").Error())
		return
	}

	workId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid work id param")
		return
	}

	var work DTO.WorkUpdate

	if err := c.BindJSON(&work); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Work.UpdateWork(userId, uint(workId), userRole, work); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h Handler) deleteWork(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	userRole, err := getUserRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	if !strings.EqualFold(userRole, models.APPLICANT) && !strings.EqualFold(userRole, models.ADMIN) {
		newErrorResponse(c, http.StatusForbidden, errors.New("not enough rights").Error())
		return
	}

	workId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid work id param")
		return
	}

	if err := h.services.Work.DeleteWork(userId, uint(workId), userRole); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
