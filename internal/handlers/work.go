package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
