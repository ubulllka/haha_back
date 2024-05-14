package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"haha/internal/models"
	"net/http"
	"strings"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(models.AuthorizationHeader)
	if header == "" {
		h.logg.Error("empty auth header")
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		h.logg.Error("invalid auth header")
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		h.logg.Error("token is empty")
		newErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	user, err := h.services.Authorization.GetUser(userId)
	if err != nil {
		h.logg.Error(err)
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(models.UserIdCtx, userId)
	c.Set(models.UserRoleCtx, user.Role)

}

func (h *Handler) GetUserId(c *gin.Context) (uint, error) {
	id, ok := c.Get(models.UserIdCtx)
	if !ok {
		h.logg.Error("user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(uint)
	if !ok {
		h.logg.Error("user id is of invalid type")
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}

func (h *Handler) GetUserRole(c *gin.Context) (string, error) {

	role, ok := c.Get(models.UserRoleCtx)
	if !ok {
		h.logg.Error("user role not found")
		return "", errors.New("user role not found")
	}

	roleStr := role.(string)
	if !strings.EqualFold(roleStr, models.ANON) && !strings.EqualFold(roleStr, models.APPLICANT) &&
		!strings.EqualFold(roleStr, models.EMPLOYER) && !strings.EqualFold(roleStr, models.ADMIN) {
		h.logg.Error("user role is of invalid type")
		return "", errors.New("user role is of invalid type")
	}

	return roleStr, nil
}
