package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/0xMarvell/go-jwt-auth-implementation/pkg/config"
	"github.com/0xMarvell/go-jwt-auth-implementation/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// Signup allows a user to register new account with the expected user details
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
			"error": "Failed to hash password",
		})
		return
	}
	// Create user object using GORM
	user := models.User{Name: signupPayload.Name, Email: signupPayload.Email, Password: string(hash)}
	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user: an account with that email already exists",
		})
		return
	}
	// Return JSON response to confirm successful creation of user
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "New user was successfully created",
	})
}

// Login allows existing user to login to the API
func Login(c *gin.Context) {
	// Get needed details (email,password) off request body
	var loginPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if c.Bind(&loginPayload) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})
		return
	}
	// Using GORM, query database to find user details
	var user models.User
	config.DB.First(&user, "email = ?", loginPayload.Email)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found. Invalid email or password",
		})
		return
	}
	// Compare password gotten off request body to user password hash stored in database
	pwdErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginPayload.Password))
	if pwdErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	// Generate JWT token
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"subject": user.ID,
		/* Expiration time on line 92 == 30 days.
		   This is jut an example implementation
		   For production, 30 days would be too much so a shorter time
		   would be more optimal.
		*/
		"expiration_time": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	// the secret key will be created by you (it can be a random sequence of characters e.g. 3r4jgnirbg8rhg08gvi0pvhh8)
	// Tip: DO NOT HARD CODE YOUR SECRET KEY
	// Store it as an environment variable instead
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to generate JWT token",
		})
		return
	}
	// Store JWT token inside httpOnly cookie (for security purposes)
	// Avoid storing your token in localstorage because it
	// becomes vulnerable to Cross-Site-Scripting (XSS) attack
	var secure, httpOnly bool = false, true // in production, secure should be set to true
	var maxAge int = 3600 * 24 * 30         // maxAge is the amount of time (IN SECONDS) the cookie will be valid for
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("auth_token", tokenString, maxAge, "", "", secure, httpOnly)
	// Return JSON Response to confirm successful storage of JWT token
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "JWT token generated successfully",
	})
}

// GetUSerDetails retireves an existing user's account details
func GetUserDetails(c *gin.Context) {
	// Retireve user details attached to request after passing through middleware
	user, _ := c.Get("user")
	// Return user details as JSON response
	c.JSON(http.StatusOK, gin.H{
		"success":      true,
		"user_details": user,
	})
}
