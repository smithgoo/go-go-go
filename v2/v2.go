package v2

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	models2 "go-go-go/models"
	"go-go-go/models"
	"go-go-go/database"
	"log"
	"net/url"
	"strings"

	//"strconv"

	"net/http"
	//"strconv"
	//"encoding/json"
)

func PingPong(c *gin.Context) {
	c.JSON(200,gin.H{
		"message":"pong",
	})
}

func GetAllInfoList(c *gin.Context)  {
	var videos []models2.VideoInfo

	paginatedResult, err := database.Paginate(c, database.DB, &models2.VideoInfo{}, &videos)
	if err != nil {
		log.Println("Error fetching videos:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	//c.JSON(http.StatusOK, paginatedResult.Data)
	c.HTML(http.StatusOK,"videoInfoList.html",gin.H{
		"products":paginatedResult.Data,
	})

}

func SearchVideoHtml(c *gin.Context)  {
	c.HTML(http.StatusOK,"videoSearch.html",gin.H{
		"title":"videoSearch",
	})
}

func SearchVideoByName(c *gin.Context)  {
	var products []models2.VideoInfo
	keyword := c.PostForm("keyword")

	query := database.DB
	if keyword != "" {
		keywordPattern := "%" + keyword + "%"
		query = query.Where(
			database.DB.Where("address LIKE ?", keywordPattern).
				Or("content LIKE ?", keywordPattern).
				Or("title LIKE ?", keywordPattern),
		)
	}

	paginatedResult, err := database.PaginateSearch(c, database.DB, &models.Product{}, &products,query)
	if err != nil {
		log.Println("Error fetching products:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, paginatedResult)

}


func PlayActionClick(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	for key, value := range queryParams {
		//log.Printf("Key: %s, Value: %s\n", key, value)
		if key == "Address" {
			queryParams["Address"] =  strings.Split(value[0],"#")
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
	videoURL := c.DefaultQuery("videoURL", "https://v6.fentvoss.com/sdv6/202403/04/hrCAjCE4P32/video/index.m3u8")

	c.HTML(http.StatusOK, "videoPlayer.html", gin.H{
		"Title":    title,
		"VideoURL": videoURL,
	})
}

