package app

import (
	"toko-niko/app/controllers"
)

func (server *Server) InitRoutes() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}
