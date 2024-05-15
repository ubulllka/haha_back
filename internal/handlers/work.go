package handlers

import (
	"github.com/gin-gonic/gin"
	"haha/internal/models"
	"net/http"
	"strconv"
	"strings"
)

func (h Handler) getListWork(c *gin.Context) {
	resumeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusBadRequest, "invalid work id param")
		return
	}

	works, err := h.services.Work.GetListWork(uint(resumeId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, works)
}

func (h Handler) deleteWork(c *gin.Context) {
	userId, err := h.GetUserId(c)
	if err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userRole, err := h.GetUserRole(c)
	if err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if !strings.EqualFold(userRole, models.APPLICANT) && !strings.EqualFold(userRole, models.ADMIN) {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusForbidden, "not enough rights")
		return
	}

	workId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusBadRequest, "invalid work id param")
		return
	}

	if err := h.services.Work.DeleteWork(userId, uint(workId), userRole); err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
