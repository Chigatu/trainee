package handler

import (
	"github.com/chigatu/service/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	segment := router.Group("/slug")
	{
		segmentId := segment.Group("/segmentId")
		{
			segmentId.POST("/", h.CreateSegment)
			segmentId.DELETE("/", h.CreateSegment)
		}

		userSegment := router.Group("userSegment")
		{
			userSegment.POST("/userSegment")
			userSegment.GET("/userSegment")
		}
	}
	return router
}
