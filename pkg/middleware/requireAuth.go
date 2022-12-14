package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/0xMarvell/go-jwt-auth-implementation/pkg/config"
	"github.com/0xMarvell/go-jwt-auth-implementation/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// RequireAuth acts as middleware to validate the received JWT token
// before granting authorization to a user
func RequireAuth(c *gin.Context) {
	// Get the cookie off the request
	tokenString, cookieErr := c.Cookie("auth_token")
	if cookieErr != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Failed to get cookie off request",
		})
	}
	// Decode/Validate JWT token stored in the cookie
	// jwt.Parse takes the token string and a function for looking up the key
	token, tokenErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if tokenErr != nil {
		log.Println("Failed to parse JWT token: ", tokenErr)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check if token has not surpassed it expiry time
		if float64(time.Now().Unix()) > claims["expiration_time"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "JWT token is expired",
			})
		}
		// Find the user using token subject
		var user models.User
		config.DB.First(&user, claims["subject"])
		if user.ID == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Failed to find user in database",
			})
		}
		// Attach user to the request
		c.Set("user", user)
		// Authorize user and continue
		c.Next() // this sends the request from the middleware to the expected controller
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Failed to validate JWT token",
		})
	}

}
