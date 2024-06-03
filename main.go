package main

import (
	"github.com/gin-gonic/gin"
	//v1 "go-go-go/v1"
	"go-go-go/database"
	"go-go-go/router"
	"go-go-go/models"
	)


func main() {

	database.InitDB()

	database.DB.AutoMigrate(&models.Product{})

	r := gin.Default()
	router.InitRouter(r)
	r.Run()

}


