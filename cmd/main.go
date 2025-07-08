package main

import (
	"database/sql"
	"go-http-cleanArch/internal/handlers"
	"go-http-cleanArch/internal/middlewares"
	"go-http-cleanArch/internal/repository"
	"go-http-cleanArch/internal/services"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://localhost/go_http_gin_db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return
	}

	logFile, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)

	repo := repository.InitRepo(db)

	serv := services.InitService(&repo)
	hand := handlers.InitHandler(&serv)

	r := gin.Default()

	r.Use(middlewares.RequestIDMiddleware())
	r.Use(middlewares.LoggerMiddlware())

	r.POST("/adduser", hand.AddUser)
	r.GET("/allusers", hand.GetAllUsers)
	r.GET("/user/:id", hand.GetUserByID)
	r.DELETE("/user/:id", hand.DelUserById)
	r.PUT("/user/:id", hand.UpdUserById)

	r.Run()
}
