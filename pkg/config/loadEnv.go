package config

import (
	"github.com/0xMarvell/go-jwt-auth-implementation/pkg/utils"
	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables
func LoadEnv() {
	dotEnvErr := godotenv.Load()
	utils.CheckErr("error loading env variables: ", dotEnvErr)
}
