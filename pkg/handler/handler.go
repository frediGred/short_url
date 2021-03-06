package handler

import (
	"api/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("create", h.create)
		api.GET("geturl/:short_url", h.geturl)
	}

	return router
}
