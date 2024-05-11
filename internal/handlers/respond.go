package handlers

import (
	"github.com/gin-gonic/gin"
	"haha/internal/models"
	"haha/internal/models/DTO"
	"net/http"
	"strconv"
)

func (h *Handler) createRespond(c *gin.Context) {
	var input DTO.RespondModel
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userRole, err := getUserRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if err := h.services.CreateRespond(userRole, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) getMyRespond(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid page param")
		return
	}

	userRole, _ := getUserRole(c)
	userId, _ := getUserId(c)

	switch userRole {
	case models.APPLICANT:
		list, pag, err := h.services.Respond.GetMyRespondAppl(userId, int64(page))

		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"list": list,
			"pag":  pag,
		})

	case models.EMPLOYER:
		list, pag, err := h.services.Respond.GetMyRespondEmpl(userId, int64(page))

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

func (h *Handler) getOtherRespond(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid page param")
		return
	}

	userRole, _ := getUserRole(c)
	userId, _ := getUserId(c)

	switch userRole {
	case models.APPLICANT:
		list, pag, err := h.services.Respond.GetOtherRespondAppl(userId, int64(page))

		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"list": list,
			"pag":  pag,
		})

	case models.EMPLOYER:
		list, pag, err := h.services.Respond.GetOtherRespondEmpl(userId, int64(page))

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
