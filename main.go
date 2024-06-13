package main

import (
	"github.com/gin-gonic/gin"
	models2 "go-go-go/models"
	"go-go-go/models"
	//v1 "go-go-go/v1"
	"go-go-go/database"
	"go-go-go/router"

	)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization, Accept, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}



func main() {

	database.InitDB( "root", "Jbnb123456", "127.0.0.1:3306","DBInfo")
	//database.InitDB( "root", "Jbnb123456", "127.0.0.1:3306","GODB")

	database.DB.AutoMigrate(&models.Product{},&models2.VideoInfo{},&models2.User{})

	r := gin.Default()
	//r.Use(gin.Recovery())
	// 设置静态文件服务，用于提供 Vue 前端打包后的静态资源
	//router.Static("/static", "./frontend/dist/static")
	r.Use(CORSMiddleware())
	router.InitRouter(r)
	r.Run(":9000")

}


