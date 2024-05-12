package handlers

import (
	"github.com/gin-gonic/gin"
	"haha/internal/models"
	"haha/internal/models/DTO"
	"net/http"
	"strconv"
)

func (h *Handler) getAllUser(c *gin.Context) {
	users, err := h.services.User.GetAllUsers()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *Handler) getInfo(c *gin.Context) {
	id, _ := getUserId(c)

	user, err := h.services.User.GetUser(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
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
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) getUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	user, err := h.services.User.GetUser(uint(id))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) isUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	userId, _ := getUserId(c)

	c.JSON(http.StatusOK, statusResponse{strconv.FormatBool(userId == uint(id))})

}

func (h *Handler) getMyListPag(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid page param")
		return
	}

	userRole, _ := getUserRole(c)
	userId, _ := getUserId(c)

	switch userRole {
	case models.APPLICANT:
		list, pag, err := h.services.Resume.GetApplAllResumesPag(userId, int64(page))

		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"list": list,
			"pag":  pag,
		})

	case models.EMPLOYER:
		list, pag, err := h.services.Vacancy.GetEmplAllVacanciesPag(userId, int64(page))

		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"list": list,
			"pag":  pag,
		})
	}
}

func (h *Handler) getListPag(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid page param")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	user, err := h.services.User.GetUser(uint(id))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userRole := user.Role

	switch userRole {
	case models.APPLICANT:
		list, pag, err := h.services.Resume.GetApplAllResumesPag(uint(id), int64(page))

		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"list": list,
			"pag":  pag,
		})

	case models.EMPLOYER:
		list, pag, err := h.services.Vacancy.GetEmplAllVacanciesPag(uint(id), int64(page))

		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"list": list,
			"pag":  pag,
		})
	}
}

func (h *Handler) getList(c *gin.Context) {
	userRole, _ := getUserRole(c)
	userId, _ := getUserId(c)

	switch userRole {
	case models.APPLICANT:
		list, err := h.services.Resume.GetApplAllResumes(userId)

		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, list)

	case models.EMPLOYER:
		list, err := h.services.Vacancy.GetEmplAllVacancies(userId)

		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, list)
	}

}
