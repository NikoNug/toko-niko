package app

import (
	"toko-niko/app/controllers"

	"github.com/gorilla/mux"
)

func (server *Server) InitRoutes() {
	server.Router = mux.NewRouter()

	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}
