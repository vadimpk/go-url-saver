package handler

import (
	"github.com/gin-gonic/gin"
	"go-urlsaver"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input go_url_saver.User
	if err := c.BindJSON(&input); err != nil {
		newError(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		newError(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input go_url_saver.User
	if err := c.BindJSON(&input); err != nil {
		newError(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newError(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
