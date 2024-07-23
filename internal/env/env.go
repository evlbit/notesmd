package env

import (
	"fmt"
	"log"
	"os"

	"github.com/lpernett/godotenv"
)

type EnvVars struct {
	DBAddress  string
	DBUser     string
	DBPassword string
	DBName     string
}

var Vars *EnvVars

func InitEnv(appEnv string) error {
	err := godotenv.Load(fmt.Sprintf(".env.%s", appEnv))
	if err != nil {
		return err
	}

	Vars = &EnvVars{
		DBAddress:  getEnv("DB_ADDRESS"),
		DBUser:     getEnv("DB_USER"),
		DBPassword: getEnv("DB_PASSWORD"),
		DBName:     getEnv("DB_NAME"),
	}

	return nil
}

func getEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("Could not find env variable %s", key)
	}
	return value
}
