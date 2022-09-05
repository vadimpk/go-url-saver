package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userID"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)

	if header == "" {
		newError(c, http.StatusUnauthorized, "empty authorization header")
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		newError(c, http.StatusUnauthorized, "invalid authorization header")
		return
	}

	userID, err := h.service.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newError(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userID)
}

func (h *Handler) getUserID(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newError(c, http.StatusInternalServerError, "user id not found")
		return -1, errors.New("user id not found")
	}

	idToInt, ok := id.(int)
	if !ok {
		newError(c, http.StatusInternalServerError, "user id not found")
		return -1, errors.New("user id not found")
	}

	return idToInt, nil
}
