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

func (h *Handler) getAllResumes(c *gin.Context) {
	resumes, err := h.services.Resume.GetAllResumes()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, resumes)
}

func (h *Handler) searchResumes(c *gin.Context) {
	q := c.Query("q")
	resumes, err := h.services.Resume.SearchResumes(q)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, resumes)
}

func (h *Handler) getApplAllResumes(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	userRole, err := getUserRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	if !strings.EqualFold(userRole, models.APPLICANT) {
		newErrorResponse(c, http.StatusForbidden, errors.New("not enough rights").Error())
		return
	}

	resumes, err := h.services.Resume.GetApplAllResumes(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, resumes)
}

func (h *Handler) getResume(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	resume, err := h.services.Resume.GetResume(uint(id))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, resume)
}

func (h *Handler) createResume(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	userRole, err := getUserRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	if !strings.EqualFold(userRole, models.APPLICANT) {
		newErrorResponse(c, http.StatusForbidden, errors.New("not enough rights").Error())
		return
	}

	var resume DTO.ResumeCreate

	if err := c.BindJSON(&resume); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Resume.CreateResume(userId, resume)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) updateResume(c *gin.Context) {
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
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var resume DTO.ResumeUpdate

	if err := c.BindJSON(&resume); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Resume.UpdateResume(userId, uint(resumeId), userRole, resume); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteResume(c *gin.Context) {
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
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := h.services.Resume.DeleteResume(userId, uint(resumeId), userRole); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
