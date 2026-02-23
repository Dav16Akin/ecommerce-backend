package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	model "github.com/Dav16Akin/ecommerce-rest-backend/internal/models"
	"github.com/Dav16Akin/ecommerce-rest-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	UpdatePassword(c *gin.Context)
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
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	user, err := h.service.GetUser(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	resp := model.UserResponse{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}

	c.JSON(http.StatusOK, resp)
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	existingUser, err := h.service.GetUser(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	var req model.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Name != nil {
		existingUser.Name = *req.Name
	}

	if req.Email != nil {
		existingUser.Email = *req.Email
	}

	if req.PhoneNumber != nil {
		existingUser.PhoneNumber = *req.PhoneNumber
	}

	if err := h.service.UpdateUser(existingUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := model.UserResponse{
		ID:          existingUser.ID,
		Name:        existingUser.Name,
		Email:       existingUser.Email,
		PhoneNumber: existingUser.PhoneNumber,
	}

	c.JSON(http.StatusOK, resp)
}

func (h *userHandler) DeleteUser(c *gin.Context) {
	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	err = h.service.DeleteUser(uint(id))
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		}
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *userHandler) UpdatePassword(c *gin.Context) {
	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var req model.UpdatePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdatePassword(uint(id), req.OldPassword, req.NewPassword); err != nil {
		switch err {
		case service.ErrUserNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		case service.ErrInvalidPassword:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid current password"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.Status(http.StatusOK)
}
