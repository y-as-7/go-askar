package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"order-system/config"
	"order-system/app/dto"
	"order-system/app/middleware"
	"order-system/app/models"
)

type AuthController struct{}

// Register new user
func (ctrl *AuthController) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	// Check if user already exists
	var existingUser models.User
	if err := config.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, dto.Response{
			Success: false,
			Error:   "User with this email already exists",
		})
		return
	}

	// Create new user
	user := models.User{
		Name:  req.Name,
		Email: req.Email,
	}

	if err := user.HashPassword(req.Password); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Error:   "Failed to hash password",
		})
		return
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Error:   "Failed to create user",
		})
		return
	}

	// Generate token
	token, err := middleware.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Error:   "Failed to generate token",
		})
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "User registered successfully",
		Data: dto.LoginResponse{
			Token: token,
			User:  user,
		},
	})
}

// Login user
func (ctrl *AuthController) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	// Find user
	var user models.User
	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Error:   "Invalid credentials",
		})
		return
	}

	// Check password
	if err := user.CheckPassword(req.Password); err != nil {
		c.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Error:   "Invalid credentials",
		})
		return
	}

	// Check if user is active
	if !user.IsActive {
		c.JSON(http.StatusForbidden, dto.Response{
			Success: false,
			Error:   "Account is inactive",
		})
		return
	}

	// Generate token
	token, err := middleware.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Error:   "Failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Login successful",
		Data: dto.LoginResponse{
			Token: token,
			User:  user,
		},
	})
}

// GetProfile returns current user profile
func (ctrl *AuthController) GetProfile(c *gin.Context) {
	userID, _ := c.Get("user_id")
	
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, dto.Response{
			Success: false,
			Error:   "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data:    user,
	})
}
