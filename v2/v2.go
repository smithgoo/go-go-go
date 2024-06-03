package v2

import (
	"github.com/gin-gonic/gin"
	"go-go-go/videoModels"
	"go-go-go/database"
	"log"
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

func SearchVideoHtml(c *gin.Context)  {
	c.HTML(http.StatusOK,"videoSearch.html",gin.H{
		"title":"videoSearch",
	})
}

func SearchVideoByName(c *gin.Context)  {
	var products []videoModels.VideoInfo
	keyword := c.PostForm("keyword")

	log.Println("Search keyword:", keyword)
	log.Println(products)
	if keyword != "" {
		query := "%" + keyword + "%"
		log.Println("Search query:", query)
		database.DB.Where("title LIKE ?", query).Find(&products)
	} else {
		database.DB.Find(&products)
	}

	c.HTML(http.StatusOK, "videoSearch.html", gin.H{
		"products": products,
	})
}