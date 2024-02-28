package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize() {
	fmt.Println("Welcome to Toko Niko")

	server.Router = mux.NewRouter()
	server.InitRoutes()
}

func (server *Server) Run(address string) {
	fmt.Println("Listening to port " + address)
	log.Fatal(http.ListenAndServe(address, server.Router))
}

func Run() {
	server := Server{}

	server.Initialize()
	server.Run(":9000")
}
