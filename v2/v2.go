package v2

import (
	//"fmt"
	"go-go-go/database"
	"go-go-go/models"
	models2 "go-go-go/models"
	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	//"strconv"

	"net/http"
	//"strconv"
	//"encoding/json"
)

func PingPong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func GetAllInfoList(c *gin.Context) {
	var videos []models2.VideoInfo

	paginatedResult, err := database.Paginate(c, database.DB, &models2.VideoInfo{}, &videos)
	if err != nil {
		log.Println("Error fetching videos:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	//c.JSON(http.StatusOK, paginatedResult.Data)
	c.HTML(http.StatusOK, "videoInfoList.html", gin.H{
		"products": paginatedResult.Data,
	})

}

func GetHomeAllInfoList(c *gin.Context) {
	var videos []models2.VideoInfo

	// 获取查询参数 page 和 size，默认值分别为 1 和 10
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	if err != nil || size < 1 {
		size = 10
	}

	// 计算 offset 和 limit
	offset := (page - 1) * size

	// 查询数据库并分页
	var total int64
	result := database.DB.Model(&models2.VideoInfo{}).Count(&total).Offset(offset).Limit(size).Find(&videos)
	if result.Error != nil {
		log.Println("Error fetching videos:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// 返回分页结果
	c.JSON(http.StatusOK, gin.H{
		"page":   page,
		"size":   size,
		"total":  total,  // 数据库中总条数
		"videos": videos, // 当前页的视频数据
	})
}

func SearchVideoHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "videoSearch.html", gin.H{
		"title": "videoSearch",
	})
}

func SearchVideoByName(c *gin.Context) {
	var products []models.VideoInfo
	keyword := strings.TrimSpace(c.PostForm("keyword")) // 处理空格

	log.Println("Received keyword:", keyword)

	query := database.DB

	if keyword != "" {
		// 将关键字拆解为单个字词
		keywords := strings.Split(keyword, "")
		query = query.Where("") // 初始化查询条件

		for _, word := range keywords {
			if word != "" { // 过滤空字符
				keywordPattern := "%" + word + "%"
				query = query.Or("title LIKE ?", keywordPattern)
			}
		}
	} else {
		// 如果关键字为空，直接返回空结果或所有数据
		c.JSON(http.StatusOK, gin.H{
			"data":   []models.VideoInfo{},
			"total":  0,
			"offset": 0,
			"limit":  0,
		})
		return
	}

	// 启用 SQL 语句调试
	query = query.Debug()

	// 打印最终 SQL 语句
	// log.Println("Generated SQL:", query.ToSQL())

	// 执行查询并分页
	paginatedResult, err := database.PaginateSearch(c, query, &models.VideoInfo{}, &products, query)
	if err != nil {
		log.Println("Error fetching products:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 成功返回结果
	c.JSON(http.StatusOK, paginatedResult)
}


func PlayActionClick(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	for key, value := range queryParams {
		//log.Printf("Key: %s, Value: %s\n", key, value)
		if key == "Address" {
			queryParams["Address"] = strings.Split(value[0], "#")
		}
	}
	log.Printf("******************")
	log.Printf("address: %s", queryParams["Address"])
	log.Printf("******************")
	queryParamsString := queryParams.Encode()

	// 使用 log.Printf 打印编码后的字符串
	paramsMap, err := url.ParseQuery(queryParamsString)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("paramsMap: %s", paramsMap)
	//c.JSON(http.StatusOK, paramsMap)
	//queryParamsJSON, err := json.Marshal(queryParamsString)
	//if err != nil {
	//	log.Printf("Error marshaling query params to JSON: %v", err)
	//	//http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	//	return
	//}
	//log.Printf("queryParamsJSON: %s", queryParamsJSON)
	//
	title := c.DefaultQuery("title", "Default Video Title")
	//https://v10.1080tg.com/202403/04/hrCAjCE4P32/video/index.m3u8
	videoURL := c.DefaultQuery("videoURL", queryParams["Address"][0])

	c.HTML(http.StatusOK, "videoPlayer.html", gin.H{
		"Title":    title,
		"VideoURL": videoURL,
	})
}
