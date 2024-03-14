package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shkuran/go-library-microservices/user-service/config"
	"github.com/shkuran/go-library-microservices/user-service/db"
	"github.com/shkuran/go-library-microservices/user-service/routes"
	"github.com/shkuran/go-library-microservices/user-service/user"
)

func main() {
	conf := config.LoadConfig()
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = conf.Database.Host
	}
	log.Printf("db_host: %s", host)
	port := conf.Database.Port
	dbUser := conf.Database.User
	pass := conf.Database.Password
	dbName := conf.Database.DbName
	sslMode := conf.Database.SslMode
	driverName := conf.Database.DriverName
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, dbUser, pass, dbName, sslMode)

	varDb, err := db.InitDB(driverName, connStr)
	if err != nil {
		log.Fatal(err)
		return
	}

	//db.CreateTables(varDb)

	server := gin.Default()

	userRepo := user.NewRepo(varDb)
	userHandler := user.NewHandler(userRepo)

	routes.RegisterRoutes(server, userHandler)

	err = server.Run(":8083")
	if err != nil {
		log.Fatal(err)
		return
	}
}
