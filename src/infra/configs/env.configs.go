package configs

import (
	"log"

	"github.com/joho/godotenv"
)

type ReturnEnvironments struct {
	DATABASE_URL string
	TOKEN_API    string
	SECRET_API   string
	GODEV        string
	PORT         string
}

func GetEnv() ReturnEnvironments {
	envs, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal(err)
	}
	var port string
	if len(envs["PORT"]) <= 0 {
		port = "3000"
	} else {
		port = envs["PORT"]
	}
	var output = ReturnEnvironments{
		DATABASE_URL: envs["DATABASE_URL"],
		TOKEN_API:    envs["TOKEN"],
		SECRET_API:   envs["SECRET"],
		GODEV:        envs["GODEV"],
		PORT:         port,
	}
	return output
}

func StartEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
}
