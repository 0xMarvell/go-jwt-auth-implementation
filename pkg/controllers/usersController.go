package controllers

import (
	"net/http"

	"github.com/0xMarvell/go-jwt-auth-implementation/pkg/config"
	"github.com/0xMarvell/go-jwt-auth-implementation/pkg/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// Get name, email, password off request body
	var signupPayload struct {
		Name     string `json:"name"`
		Email    string `json:"email" gorm:"unique"`
		Password string `json:"password"`
	}
	if c.Bind(&signupPayload) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})
		return
	}

	// Hash password gotten from request body
	hash, err := bcrypt.GenerateFromPassword([]byte(signupPayload.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})
		return
	}

	// Create user object using GORM
	user := models.User{Name: signupPayload.Name, Email: signupPayload.Email, Password: string(hash)}
	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})
		return
	}

	// Return a response object
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "New user was successfully created",
	})
}
