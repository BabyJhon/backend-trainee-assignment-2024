package handlers

import (
	"net/http"

	"github.com/BabyJhon/backend-trainee-assignment-2024/internal/entity"
	"github.com/BabyJhon/backend-trainee-assignment-2024/internal/middleware"
	"github.com/gin-gonic/gin"
)

func (h *Handler) register(c *gin.Context) { //регистрация
	var input entity.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	isValid, err := middleware.IsRegisterInputValid(input)
	if err!=nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	if !isValid {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Auth.CreateUser(c, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) login(c *gin.Context) {
	var input entity.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
	}

	// token, err := h.services.Authorization.GenerateToken(context.Background(), input.UserName, input.Password)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// c.JSON(http.StatusOK, map[string]interface{}{
	// 	"token": token,
	// })
}

func (h *Handler) dummyLogin(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	userType := queryParams.Get("user_type")

	isValid, err := middleware.IsDummyLoginInputValid(userType)
	if err!=nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	if !isValid {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Auth.GenerateToken(userType)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
