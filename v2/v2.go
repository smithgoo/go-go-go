package v2

import (
	"github.com/gin-gonic/gin"
	"go-go-go/videoModels"
	"go-go-go/database"
	"log"
	"net/http"
	"strconv"
)

func PingPong(c *gin.Context) {
	c.JSON(200,gin.H{
		"message":"pong",
	})
}

func GetAllInfoList(c *gin.Context)  {
	var videos []videoModels.VideoInfo

	// 获取分页参数
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	// 计算偏移量
	offset := (page - 1) * pageSize

	// 执行分页查询
	result := database.DB.Limit(pageSize).Offset(offset).Find(&videos)
	if result.Error != nil {
		log.Println("Error fetching videos:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// 获取总记录数
	var total int64
	database.DB.Model(&videoModels.VideoInfo{}).Count(&total)

	// 返回分页结果
	c.JSON(http.StatusOK, gin.H{
		"page":      page,
		"size": pageSize,
		"total":     total,
		"videos":    videos,
	})
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