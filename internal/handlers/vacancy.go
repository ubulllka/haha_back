package handlers

import (
	"github.com/gin-gonic/gin"
	"haha/internal/models"
	"haha/internal/models/DTO"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) getAllVacancies(c *gin.Context) {
	vacancies, err := h.services.Vacancy.GetAllVacancies()
	if err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, vacancies)
}

func (h *Handler) searchVacanciesAnon(c *gin.Context) {
	q := c.Query("q")

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusBadRequest, "invalid page param")
		return
	}

	vacancies, pag, err := h.services.Vacancy.SearchVacanciesAnon(int64(page), q)
	if err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"list": vacancies,
		"pag":  pag,
	})
}

func (h *Handler) getVacancyAnon(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	vacancy, err := h.services.Vacancy.GetVacancyAnon(uint(id))
	if err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, vacancy)
}

func (h *Handler) searchVacancies(c *gin.Context) {
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

	vacancies, pag, err := h.services.Vacancy.SearchVacancies(userId, int64(page), q)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"list": vacancies,
		"pag":  pag,
	})
}

func (h *Handler) getVacancy(c *gin.Context) {
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

	vacancy, err := h.services.Vacancy.GetVacancy(userId, uint(id))
	if err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, vacancy)
}

func (h *Handler) createVacancy(c *gin.Context) {
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

	if !strings.EqualFold(userRole, models.EMPLOYER) {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusForbidden, "not enough rights")
		return
	}

	var vacancy DTO.VacancyCreate

	if err := c.BindJSON(&vacancy); err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Vacancy.CreateVacancy(userId, vacancy); err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h Handler) updateVacancy(c *gin.Context) {
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

	if !strings.EqualFold(userRole, models.EMPLOYER) && !strings.EqualFold(userRole, models.ADMIN) {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusForbidden, "not enough rights")
		return
	}

	vacancyId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var vacancy DTO.VacancyUpdate

	if err := c.BindJSON(&vacancy); err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Vacancy.UpdateVacancy(userId, uint(vacancyId), userRole, vacancy); err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h Handler) deleteVacancy(c *gin.Context) {
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

	if !strings.EqualFold(userRole, models.EMPLOYER) && !strings.EqualFold(userRole, models.ADMIN) {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusForbidden, "not enough rights")
		return
	}

	vacancyId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := h.services.Vacancy.DeleteVacancy(userId, uint(vacancyId), userRole); err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
