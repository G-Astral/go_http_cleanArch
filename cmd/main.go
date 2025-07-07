package main

import (
	"database/sql"
	"go-http-cleanArch/internal/handlers"
	"go-http-cleanArch/internal/repository"
	"go-http-cleanArch/internal/services"
	"log"

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

	repo := repository.InitRepo(db)

	serv := services.InitService(&repo)
	hand := handlers.InitHandler(&serv)

	r := gin.Default()

	r.POST("/adduser", hand.AddUser)
	r.GET("/allusers", hand.GetAllUsers)

	r.Run()
}
