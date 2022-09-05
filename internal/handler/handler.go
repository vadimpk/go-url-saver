package handler

import (
	"github.com/gin-gonic/gin"
	"go-urlsaver/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api", h.userIdentity)
	{
		urls := api.Group("/urls")
		{
			urls.POST("/", h.createUrl)
			urls.GET("/", h.getUrls)
			urls.GET("/:url_id", h.getUrlByID)
			urls.PUT("/:url_id", h.updateUrl)
			urls.DELETE("/:url_id", h.deleteUrl)
		}
	}

	return router
}
