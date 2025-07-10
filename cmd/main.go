package main

import (
	"database/sql"
	"go-http-cleanArch/internal/handlers"
	"go-http-cleanArch/internal/logger"
	"go-http-cleanArch/internal/middlewares"
	"go-http-cleanArch/internal/repository"
	"go-http-cleanArch/internal/services"
	"log"
	"os"

	// "fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	var connStr string
	host := os.Getenv("DB_HOST")
	if host == "" {
		connStr = "postgres://localhost/go_http_gin_db?sslmode=disable"
	} else {
		connStr = "postgres://postgres@db:5432/go_http_gin_db?sslmode=disable"
	}

	// fmt.Println("connStr =", connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return
	}

	// row := db.QueryRow("SELECT current_user")
	// var ut string
	// err = row.Scan(&ut)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Текущий пользователь в PostgreSQL:", ut)

	logger.InitLogger()
	defer logger.Logger.Sync()

	repo := repository.InitRepo(db)

	serv := services.InitService(&repo)
	hand := handlers.InitHandler(&serv)

	r := gin.Default()

	r.Use(middlewares.RequestIDContexMiddleware())
	r.Use(middlewares.LoggerMiddlware(logger.Logger))

	r.POST("/adduser", hand.AddUser)
	r.GET("/allusers", hand.GetAllUsers)
	r.GET("/user/:id", hand.GetUserByID)
	r.DELETE("/user/:id", hand.DelUserById)
	r.PUT("/user/:id", hand.UpdUserById)

	r.Run()
}
