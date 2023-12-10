package handler

import (
	"task-level-0/internal/service"

	"github.com/gin-gonic/gin"
)

type OrderReader interface {
	GetOrder(c *gin.Context)
}

type Handler struct {
	service *service.Service
}

func NewHandler(serv *service.Service) *Handler {
	return &Handler{
		service: serv,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		order := api.Group("/order")
		{
			order.GET("/:id", h.GetOrder)
		}
	}

	return router
}
