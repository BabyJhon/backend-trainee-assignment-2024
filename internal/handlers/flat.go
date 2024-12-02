package handlers

import (
	"net/http"

	"github.com/BabyJhon/backend-trainee-assignment-2024/internal/entity"
	"github.com/BabyJhon/backend-trainee-assignment-2024/internal/middleware"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateFlat(c *gin.Context) {
	var inPut entity.Flat

	if err := c.BindJSON(&inPut); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := middleware.IsCreateFlatInputValid(inPut)
	if err!=nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	flat, err := h.services.Flat.CreateFlat(c, inPut)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":       flat.Id,
		"house_id": flat.HouseId,
		"price":    flat.Price,
		"rooms":    flat.Rooms,
		"status":   flat.Status,
	})
}
