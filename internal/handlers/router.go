package handlers

import (
	"github.com/BabyJhon/backend-trainee-assignment-2024/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/dummyLogin", h.dummyLogin) //+
	router.POST("/login", h.login) //+
	router.POST("/register", h.register) //+

	authOnly := router.Group("/", h.UserIdentity)
	{
		authOnly.GET("/house/:id", h.createHouse)
		authOnly.POST("/house/:id/subscribe", h.createHouse)
		authOnly.POST("/flat/create", h.CreateFlat) //+
	}

	moderationsOnly := router.Group("/", h.ModeratorIdentity)
	{
		moderationsOnly.POST("/house/create", h.createHouse) //+
		moderationsOnly.POST("/flat/update", h.createHouse)
	}

	return router
}
