package server

import (
	"abstract-entities/src/infra/configs"
	"abstract-entities/src/infra/ports/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func StartAPIServer() {
	envs := configs.GetEnv()
	var debug bool = false
	if envs.GODEV == "True" || envs.GODEV == "true" {
		debug = true
	}
	SettingsCors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content", "Accept", "Content-Type"},
		Debug:            debug,
	})
	mux := routes.Routes()
	port := fmt.Sprintf(":%v", envs.PORT)
	fmt.Println(" [ * ] Start server in port", port)
	log.Fatal(http.ListenAndServe(port, SettingsCors.Handler(mux)))
}
