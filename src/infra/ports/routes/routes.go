package routes

import (
	"abstract-entities/src/infra/ports/controllers"
	"net/http"
)

func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/status", controllers.HealtCheck)
	mux.HandleFunc("/", controllers.HealtCheck)
	mux.HandleFunc("/cnh", controllers.CreateCNH)
	return mux
}
