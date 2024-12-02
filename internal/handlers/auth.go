package handlers

import (
	"net/http"
	"strings"

	"github.com/BabyJhon/backend-trainee-assignment-2024/internal/entity"
	"github.com/BabyJhon/backend-trainee-assignment-2024/internal/middleware"
	"github.com/gin-gonic/gin"
)

const (
	authoriationHeader = "Authorization"
)

func (h *Handler) register(c *gin.Context) {
	var input entity.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := middleware.IsRegisterInputValid(input)
	if err != nil {
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

	err := middleware.IsLoginInputValid(input.Id, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Login(c, input.Id, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) dummyLogin(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	userType := queryParams.Get("user_type")

	err := middleware.IsDummyLoginInputValid(userType)
	if err != nil {
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

func (h *Handler) ModeratorIdentity(c *gin.Context) {
	header := c.GetHeader(authoriationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty request header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid header size")
	}

	userType, err := h.services.Auth.Parsetoken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	if userType != "moderator" {
		newErrorResponse(c, http.StatusUnauthorized, "wrong user type")
	}
}

func (h *Handler) UserIdentity(c *gin.Context) {
	header := c.GetHeader(authoriationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty request header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid header size")
	}

	userType, err := h.services.Auth.Parsetoken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	if userType != "moderator" && userType != "client" {
		newErrorResponse(c, http.StatusUnauthorized, "wrong user type")
	}
}
