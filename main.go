package main

import (
	"github.com/gin-gonic/gin"
	models2 "go-go-go/models"
	"go-go-go/models"

	//v1 "go-go-go/v1"
	"go-go-go/database"
	"go-go-go/router"
	)


func main() {

	database.InitDB( "root", "Jbnb123456", "127.0.0.1:3306","DBInfo")
	//database.InitDB( "root", "Jbnb123456", "127.0.0.1:3306","GODB")

	database.DB.AutoMigrate(&models.Product{},&models2.VideoInfo{})

	r := gin.Default()
	router.InitRouter(r)
	r.Run()

}


