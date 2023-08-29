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

	user := router.Group("/user")
	{
		user.POST("/user-id")
	}

	segment := router.Group("/slug")
	{
		segmentId := segment.Group("/segmentId")
		{
			segmentId.POST("/:id")
			segmentId.DELETE("/:id")
		}

		userSegment := router.Group("userSegment")
		{
			userSegment.GET("/userSegment")
		}
	}
	return router
}
