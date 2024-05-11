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

func (h *Handler) getAllVacancies(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid page param")
		return
	}

	vacancies, pag, err := h.services.Vacancy.GetAllVacancies(int64(page))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"list": vacancies,
		"pag":  pag,
	})
}

func (h *Handler) searchVacancies(c *gin.Context) {
	q := c.Query("q")

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid page param")
		return
	}

	vacancies, pag, err := h.services.Vacancy.SearchVacancies(int64(page), q)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"list": vacancies,
		"pag":  pag,
	})
}

func (h *Handler) getVacancy(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	vacancy, err := h.services.Vacancy.GetVacancy(uint(id))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, vacancy)
}

func (h *Handler) createVacancy(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	userRole, err := getUserRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	if !strings.EqualFold(userRole, models.EMPLOYER) {
		newErrorResponse(c, http.StatusForbidden, errors.New("not enough rights").Error())
		return
	}

	var vacancy DTO.VacancyCreate

	if err := c.BindJSON(&vacancy); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Vacancy.CreateVacancy(userId, vacancy)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h Handler) updateVacancy(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	userRole, err := getUserRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	if !strings.EqualFold(userRole, models.EMPLOYER) && !strings.EqualFold(userRole, models.ADMIN) {
		newErrorResponse(c, http.StatusForbidden, errors.New("not enough rights").Error())
		return
	}

	vacancyId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var vacancy DTO.VacancyUpdate

	if err := c.BindJSON(&vacancy); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Vacancy.UpdateVacancy(userId, uint(vacancyId), userRole, vacancy); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h Handler) deleteVacancy(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	userRole, err := getUserRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	if !strings.EqualFold(userRole, models.EMPLOYER) && !strings.EqualFold(userRole, models.ADMIN) {
		newErrorResponse(c, http.StatusForbidden, errors.New("not enough rights").Error())
		return
	}

	vacancyId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := h.services.Vacancy.DeleteVacancy(userId, uint(vacancyId), userRole); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
