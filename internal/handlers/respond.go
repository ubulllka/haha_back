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
	userId, _ := getUserId(c)
	userRole, _ := getUserRole(c)

	if err := h.services.CreateRespond(userId, userRole, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) updateRespond(c *gin.Context) {
	var input DTO.RespondUpdate
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.Status != models.ACCEPT && input.Status != models.DECLINE && input.Status != models.WAIT {
		newErrorResponse(c, http.StatusBadRequest, "invalid status param")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	userId, _ := getUserId(c)
	userRole, _ := getUserRole(c)

	if err := h.services.Respond.UpdateRespond(userId, userRole, uint(id), input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) getMyRespond(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	userRole, _ := getUserRole(c)
	userId, _ := getUserId(c)

	respond, err := h.services.Respond.GetMyRespond(userId, userRole, uint(id))

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, respond)

}

func (h *Handler) getOtherRespond(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	userRole, _ := getUserRole(c)
	userId, _ := getUserId(c)

	respond, err := h.services.Respond.GetOtherRespond(userId, userRole, uint(id))

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, respond)

}

func (h *Handler) getMyAllResponds(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid page param")
		return
	}

	filter := c.Query("filter")

	userRole, _ := getUserRole(c)
	userId, _ := getUserId(c)

	list, pag, err := h.services.Respond.GetMyAllResponds(userId, userRole, int64(page), filter)

	if err != nil {
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
		newErrorResponse(c, http.StatusBadRequest, "invalid page param")
		return
	}

	filter := c.Query("filter")

	userId, _ := getUserId(c)
	userRole, _ := getUserRole(c)

	list, pag, err := h.services.Respond.GetOtherAllResponds(userId, userRole, int64(page), filter)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"list": list,
		"pag":  pag,
	})

}

func (h *Handler) deleteMyRespond(c *gin.Context) {
	userRole, _ := getUserRole(c)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	userId, _ := getUserId(c)

	if err := h.services.Respond.DeleteMyRespond(userId, userRole, uint(id)); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})

}

func (h *Handler) deleteOtherRespond(c *gin.Context) {
	userRole, _ := getUserRole(c)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	userId, _ := getUserId(c)

	if err := h.services.Respond.DeleteOtherRespond(userId, userRole, uint(id)); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})

}
