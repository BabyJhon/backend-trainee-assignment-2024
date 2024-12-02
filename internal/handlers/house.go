package handlers

import (
	"net/http"

	"github.com/BabyJhon/backend-trainee-assignment-2024/internal/entity"
	"github.com/BabyJhon/backend-trainee-assignment-2024/internal/middleware"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createHouse(c *gin.Context) {
	var inPut entity.House

	if err := c.BindJSON(&inPut); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	isValid, err := middleware.IsCreateHouseInputValid(inPut)
	if !isValid {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	house, err := h.services.House.CreateHouse(c, inPut)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":         house.Id,
		"address":    house.Address,
		"year":       house.Year,
		"developer":  house.Developer,
		"created_at": house.CreatedAt,
		"updated_at": house.UpdatedAt,
	})

}
