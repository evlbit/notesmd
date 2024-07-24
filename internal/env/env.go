package env

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/lpernett/godotenv"
)

type EnvVars struct {
	DBAddress   string
	DBUser      string
	DBPassword  string
	DBName      string
	JWTSecret   string
	JWTExpHours int
}

var Vars *EnvVars

func InitEnv(appEnv string) error {
	err := godotenv.Load(fmt.Sprintf(".env.%s", appEnv))
	if err != nil {
		return err
	}

	Vars = &EnvVars{
		DBAddress:   getEnv("DB_ADDRESS"),
		DBUser:      getEnv("DB_USER"),
		DBPassword:  getEnv("DB_PASSWORD"),
		DBName:      getEnv("DB_NAME"),
		JWTSecret:   getEnv("JWT_SECRET"),
		JWTExpHours: getEnvInt("JWT_EXP_HOURS"),
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

func getEnvInt(key string) int {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("Could not find env variable %s", key)
	}

	valueI64, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		log.Fatalf("Could not parse env variable %s to int. Error: %s", key, err.Error())
	}

	return int(valueI64)
}
