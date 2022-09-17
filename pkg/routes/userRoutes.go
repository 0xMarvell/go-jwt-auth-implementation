package routes

import (
	"github.com/0xMarvell/go-jwt-auth-implementation/pkg/controllers"
	"github.com/0xMarvell/go-jwt-auth-implementation/pkg/middleware"
	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes handles all routing for the API
func RegisterUserRoutes(r *gin.Engine) {
	r.POST("/api/v1/user/signup", controllers.Signup)
	r.POST("/api/v1/user/login", controllers.Login)
	r.GET("/api/v1/user/details", middleware.RequireAuth, controllers.GetUserDetails)
}
