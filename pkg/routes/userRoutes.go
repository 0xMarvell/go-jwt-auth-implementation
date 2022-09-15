package routes

import (
	"github.com/0xMarvell/go-jwt-auth-implementation/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	// panic("not completed")
	r.POST("/api/v1/user/signup", controllers.Signup)
}
