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
		h.logg.Error(err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

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

	if err := h.services.CreateRespond(userId, userRole, input); err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) updateRespond(c *gin.Context) {
	var input DTO.RespondUpdate

	if err := c.BindJSON(&input); err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.Status != models.ACCEPT && input.Status != models.DECLINE && input.Status != models.WAIT {
		h.logg.Error("invalid status param")
		newErrorResponse(c, http.StatusBadRequest, "invalid status param")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logg.Error("invalid id param")
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

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

	if err := h.services.Respond.UpdateRespond(userId, userRole, uint(id), input); err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) getMyAllResponds(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusBadRequest, "invalid page param")
		return
	}

	filter := c.Query("filter")

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

	list, pag, err := h.services.Respond.GetMyAllResponds(userId, userRole, int64(page), filter)

	if err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"list": list,
		"pag":  pag,
	})

}

func (h *Handler) getOtherAllResponds(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusBadRequest, "invalid page param")
		return
	}

	filter := c.Query("filter")

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

	list, pag, err := h.services.Respond.GetOtherAllResponds(userId, userRole, int64(page), filter)

	if err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"list": list,
		"pag":  pag,
	})

}

func (h *Handler) deleteMyRespond(c *gin.Context) {
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

	userRole, err := h.GetUserRole(c)
	if err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Respond.DeleteMyRespond(userId, userRole, uint(id)); err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})

}

func (h *Handler) deleteOtherRespond(c *gin.Context) {
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

	userRole, err := h.GetUserRole(c)
	if err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Respond.DeleteOtherRespond(userId, userRole, uint(id)); err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})

}
