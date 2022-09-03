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
	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.signIn)
			lists.GET("/", h.signIn)
			lists.GET("/:id", h.signIn)
			lists.PUT("/:id", h.signIn)
			lists.DELETE("/:id", h.signIn)

			urls := api.Group("/:id/urls")
			{
				urls.POST("/", h.signIn)
				urls.GET("/", h.signIn)
				urls.GET("/:url_id", h.signIn)
				urls.PUT("/:url_id", h.signIn)
				urls.DELETE("/:url_id", h.signIn)
			}
		}
	}

	return router
}
