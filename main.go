package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type AppConfig struct {
	DBHost         string `mapstructure:"DB_HOST"`
	DBPort         int    `mapstructure:"DB_PORT"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPassword     string `mapstructure:"DB_PASSWORD"`
	DBName         string `mapstructure:"DB_NAME"`
	HttpServerPort int    `mapstructure:"HTTP_SERVER_PORT"`
}

// cfg is the global configuration object for app
var cfg AppConfig

// db is the global db instance for app
var db *sql.DB

func main() {
	err := LoadConfig(".")
	if err != nil {
		panic(err)
	}
	fmt.Println("Config loaded!")

	err = ConnectDB()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to DB!")

	http.HandleFunc("/ingest", handlerFunc)
	portNum := fmt.Sprintf(":%d", cfg.HttpServerPort)
	err = http.ListenAndServe(portNum, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listening on port", portNum)
}
