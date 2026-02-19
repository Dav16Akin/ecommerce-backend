package handler

import (
	"net/http"
	"strings"

	model "github.com/Dav16Akin/ecommerce-rest-backend/internal/models"
	"github.com/Dav16Akin/ecommerce-rest-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetUser(c *gin.Context)
	CreateUser(c *gin.Context)
}

type userHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return &userHandler{service: service}
}

func (h *userHandler) CreateUser(c *gin.Context) {
	var req model.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Name:        strings.TrimSpace(req.Name),
		Email:       strings.TrimSpace(req.Email),
		PhoneNumber: strings.TrimSpace(req.PhoneNumber),
		Password:    req.Password,
	}

	if err := h.service.CreateUser(&user); err != nil {

		if strings.Contains(err.Error(), "required") {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	resp := model.UserResponse{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *userHandler) GetUser(c *gin.Context) {
	user, err := h.service.GetUser(1)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	c.JSON(200, user)
}
