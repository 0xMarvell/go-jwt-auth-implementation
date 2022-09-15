package main

import (
	"log"

	"github.com/0xMarvell/go-jwt-auth-implementation/pkg/config"
)

func init() {
	config.LoadEnv()
	config.Connect()
}

func main() {
	log.Println("Hello world")
}
