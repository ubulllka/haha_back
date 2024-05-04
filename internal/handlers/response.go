package handlers

import (
	"github.com/gin-gonic/gin"
	"haha/internal/logger"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logg := logger.GetLogger()
	logg.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
