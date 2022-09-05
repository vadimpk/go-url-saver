package handler

import (
	"github.com/gin-gonic/gin"
	"go-urlsaver"
	"net/http"
	"strconv"
)

func (h *Handler) createUrl(c *gin.Context) {

	userID, err := h.getUserID(c)
	if err != nil {
		return
	}

	var input go_url_saver.Url
	err = c.BindJSON(&input)
	if err != nil {
		newError(c, http.StatusBadRequest, err.Error())
		return
	}

	urlID, err := h.service.URL.CreateURL(userID, input)
	if err != nil {
		newError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"url_id": urlID,
	})

}

type getAllUrlsResponse struct {
	Data []go_url_saver.UrlResponse `json:"data"`
}

func (h *Handler) getUrls(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		return
	}

	urls, err := h.service.URL.GetAll(userID)
	if err != nil {
		newError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllUrlsResponse{
		Data: urls,
	})
}

func (h *Handler) getUrlByID(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		return
	}

	urlID, err := strconv.Atoi(c.Param("url_id"))
	if err != nil {
		newError(c, http.StatusBadRequest, err.Error())
		return
	}

	url, err := h.service.URL.GetByID(userID, urlID)
	if err != nil {
		newError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, url)
}

func (h *Handler) updateUrl(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		return
	}

	urlID, err := strconv.Atoi(c.Param("url_id"))
	if err != nil {
		newError(c, http.StatusBadRequest, err.Error())
		return
	}

	var input go_url_saver.UpdateUrl
	if err := c.BindJSON(&input); err != nil {
		newError(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.URL.UpdateURL(userID, urlID, input)
	if err != nil {
		newError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "updated",
	})

}

func (h *Handler) deleteUrl(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		return
	}

	urlID, err := strconv.Atoi(c.Param("url_id"))
	if err != nil {
		newError(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.URL.DeleteURL(userID, urlID)
	if err != nil {
		newError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "deleted",
	})
}
