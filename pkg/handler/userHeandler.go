package handler

import (
	"net/http"

	"github.com/chigatu/service/entity"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ManageUserToSegments(c *gin.Context) {
	var input entity.ManageUserToSegmentsRequest
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	resp, err := h.services.ManageUserToSegments(input.SlugsToAdd, input.SlugsToRemove, input.UserId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, entity.ManageUserToSegmentsResponse{
		SlugsHaveBeenAdded:   resp.SlugsHaveBeenAdded,
		SlugsHaveBeenRemoved: resp.SlugsHaveBeenRemoved,
		UserId:               resp.UserId,
	})
}

func (h *Handler) GetUserSegments(c *gin.Context) {
	var input entity.GetUserSegmentsRequest
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	slugs, err := h.services.User.GetUserSegments(input.UserId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, entity.GetUserSegmentsResponse{Slugs: slugs})
}
