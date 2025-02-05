package handlers

import (
	"net/http"
	"strconv"

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

	err := middleware.IsCreateHouseInputValid(inPut)
	if err != nil {
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

func (h *Handler) GetFlatsByHouse(c *gin.Context) {

	houseId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = middleware.IsGetHousesInputValid(houseId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userType, exists := c.Get("userType")
	if exists != true {
		newErrorResponse(c, http.StatusUnauthorized, "user type not exists")
		return
	}

	flats, err := h.services.House.GetFlatsByHouse(c, houseId, userType.(string))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, &flats)
}
