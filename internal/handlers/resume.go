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

var (
	errAuth = errors.New("not enough rights")
)

func (h *Handler) getAllResumes(c *gin.Context) {
	resumes, err := h.services.Resume.GetAllResumes()
	if err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, resumes)
}

func (h *Handler) searchResumesAnon(c *gin.Context) {
	q := c.Query("q")

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusBadRequest, "invalid page param")
		return
	}

	resumes, pag, err := h.services.Resume.SearchResumesAnon(int64(page), q)
	if err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"list": resumes,
		"pag":  pag,
	})
}

func (h *Handler) getResumeAnon(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	resume, err := h.services.Resume.GetResumeAnon(uint(id))
	if err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, resume)
}

func (h *Handler) searchResumes(c *gin.Context) {
	q := c.Query("q")

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusBadRequest, "invalid page param")
		return
	}

	userId, err := h.GetUserId(c)
	if err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	resumes, pag, err := h.services.Resume.SearchResumes(userId, int64(page), q)
	if err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"list": resumes,
		"pag":  pag,
	})
}

func (h *Handler) getResume(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	userId, err := h.GetUserId(c)
	if err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	resume, err := h.services.Resume.GetResume(userId, uint(id))
	if err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, resume)
}

func (h *Handler) createResume(c *gin.Context) {
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

	if !strings.EqualFold(userRole, models.APPLICANT) {
		h.logg.Error(errAuth)
		newErrorResponse(c, http.StatusForbidden, "not enough rights")
		return
	}

	var resume DTO.ResumeCreate

	if err := c.BindJSON(&resume); err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Resume.CreateResume(userId, resume); err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) updateResume(c *gin.Context) {
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
		h.logg.Error(errAuth)
		newErrorResponse(c, http.StatusForbidden, "not enough rights")
		return
	}

	resumeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logg.Error("invalid id param")
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var resume DTO.ResumeUpdate

	if err := c.BindJSON(&resume); err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Resume.UpdateResume(userId, uint(resumeId), userRole, resume); err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteResume(c *gin.Context) {
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
		h.logg.Error(errAuth)
		newErrorResponse(c, http.StatusForbidden, "not enough rights")
		return
	}

	resumeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := h.services.Resume.DeleteResume(userId, uint(resumeId), userRole); err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
