package main

import (
	"log"

	"github.com/0xMarvell/go-jwt-auth-implementation/pkg/config"
	"github.com/0xMarvell/go-jwt-auth-implementation/pkg/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.Connect()
	config.RunMigrations()
}

func main() {
	// log.Println("Hello world")
	r := gin.Default()
	routes.RegisterUserRoutes(r)

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	log.Fatal(r.Run())
}
