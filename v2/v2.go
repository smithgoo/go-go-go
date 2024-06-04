package v2

import (
	//"fmt" 
	"github.com/gin-gonic/gin"
	"go-go-go/models"
	"go-go-go/videoModels"
	"go-go-go/database"
	"log"
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
	var videos []videoModels.VideoInfo

	paginatedResult, err := database.Paginate(c, database.DB, &videoModels.VideoInfo{}, &videos)
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
	var products []videoModels.VideoInfo
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
		log.Printf("Key: %s, Value: %s\n", key, value)
	}

	//c.JSON(http.StatusOK, gin.H{
	//	"message": "Parameters received",
	//	"params":  queryParams,
	//})

	c.HTML(http.StatusOK, "videoPlayer.html", gin.H{
		"info": queryParams,
	})
}

