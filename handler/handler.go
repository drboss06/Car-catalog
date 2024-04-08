package handler

import (
	"carDirectory/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.GET("/cars", h.GetCars)
		api.DELETE("/cars/:id", h.DeleteCar)
		api.PUT("/cars/:id", h.UpdateCar)
		api.POST("/cars", h.AddCar)
	}

	return router
}
