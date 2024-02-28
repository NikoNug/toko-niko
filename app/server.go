package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

func (server *Server) Initialize(appConfig AppConfig) {
	fmt.Println("Welcome to " + appConfig.AppName)

	server.Router = mux.NewRouter()
	server.InitRoutes()
}

func (server *Server) Run(address string) {
	fmt.Println("Listening to port " + address)
	log.Fatal(http.ListenAndServe(address, server.Router))
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func Run() {
	server := Server{}
	appConfig := AppConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	appConfig.AppName = getEnv("APP_NAME", "Toko Niko Backup")
	appConfig.AppEnv = getEnv("APP_ENV", "development")
	appConfig.AppPort = getEnv("APP_PORT", "9000")

	server.Initialize(appConfig)
	server.Run(":" + appConfig.AppPort)
}
