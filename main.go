package main

import (
	"github.com/gin-gonic/gin"
	models2 "go-go-go/models"
	"go-go-go/models"

	//v1 "go-go-go/v1"
	"go-go-go/database"
	"go-go-go/router"
	"fmt"
	"time"
	)


func main() {

	database.InitDB( "root", "Jbnb123456", "127.0.0.1:3306","DBInfo")
	//database.InitDB( "root", "Jbnb123456", "127.0.0.1:3306","GODB")

	database.DB.AutoMigrate(&models.Product{},&models2.VideoInfo{})

	r := gin.Default()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// 你的自定义格式
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	//r.Use(gin.Recovery())
	router.InitRouter(r)
	r.Run(":9000")

}


