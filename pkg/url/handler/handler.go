package handler

import (
	"github.com/gin-gonic/gin"

	service "toz/pkg/url/service"
)

type HandlerUrl struct {
	services *service.ServiceUrl
}

func NewHandlerUrl(services *service.ServiceUrl) *HandlerUrl {
	return &HandlerUrl{services: services}
}

func (h *HandlerUrl) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/get/:url", h.GetFull)
	router.POST("/create", h.Create)

	return router
}
