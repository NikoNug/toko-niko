package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
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

type DBConfig struct {
	DBHost string
	DBUser string
	DBPass string
	DBPort string
	DBName string
}

func (server *Server) Initialize(appConfig AppConfig, dbConfig DBConfig) {
	fmt.Println("Welcome to " + appConfig.AppName)

	server.InitDB(dbConfig)
	server.InitRoutes()
}

func (server *Server) Run(address string) {
	fmt.Println("Listening to port " + address)
	log.Fatal(http.ListenAndServe(address, server.Router))
}

func (server *Server) InitDB(dbConfig DBConfig) {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbConfig.DBUser, dbConfig.DBPass, dbConfig.DBHost, dbConfig.DBPort, dbConfig.DBName)
	server.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to Database")
	}

	for _, model := range RegisterModels() {
		err = server.DB.Debug().AutoMigrate(model.Model)

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Migration Database Successful")
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
	dbConfig := DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	appConfig.AppName = getEnv("APP_NAME", "Toko Niko Backup")
	appConfig.AppEnv = getEnv("APP_ENV", "development")
	appConfig.AppPort = getEnv("APP_PORT", "9000")

	dbConfig.DBHost = getEnv("DB_HOST", "localhost")
	dbConfig.DBUser = getEnv("DB_USER", "root")
	dbConfig.DBPass = getEnv("DB_PASS", "")
	dbConfig.DBPort = getEnv("DB_PORT", "3306")
	dbConfig.DBName = getEnv("DB_NAME", "gotoko")

	server.Initialize(appConfig, dbConfig)
	server.Run(":" + appConfig.AppPort)
}
