package v2

import (
	"github.com/gin-gonic/gin"
	"go-go-go/videoModels"
	"go-go-go/database"
	"net/http"
)

func PingPong(c *gin.Context) {
	c.JSON(200,gin.H{
		"message":"pong",
	})
}

func GetAllInfoList(c *gin.Context)  {
	var products []videoModels.VideoInfo
	database.DB.Find(&products)
	c.JSON(http.StatusOK,products)
}