package handler

import (
	"net/http"

	"github.com/chigatu/service/entity"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateSegment(c *gin.Context) {

	var input entity.SegmentRequest

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Segment.Create(input.Slug)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) DeleteSegment(c *gin.Context) {
	var input entity.SegmentRequest
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.services.Segment.Delete(input.Slug); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
	})
}
