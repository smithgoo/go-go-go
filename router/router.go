package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	v2 "go-go-go/v2"
	v1 "go-go-go/v1"
	"fmt"
	"time"
)
import _ "net/url"
import _ "strconv"

func InitRouter(r *gin.Engine) {

	r.LoadHTMLGlob("templates/*")
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

	GroupV1 := r.Group("/v1")
	{
		GroupV1.GET("/homePage",v1.ShowHomePage)
		GroupV1.GET("/product/add", v1.ShowAddProductPage)
		GroupV1.GET("/products",v1.ListProducts)
		GroupV1.POST("/product/add",v1.AddProduct)
		GroupV1.GET("/product/:id",v1.GetProduct)
		GroupV1.DELETE("/product/:id",v1.DeleteProduct)
		GroupV1.GET("/product/edit/:id", v1.ShowEditProductPage)
		GroupV1.POST("/product/edit/:id", v1.EditProduct)
		GroupV1.GET("/products/search",v1.SearchGetProductWithKw)
		GroupV1.POST("/products/search",v1.SearchProductWithKw)
	}

	GroupV2 := r.Group("/v2")
	{
		GroupV2.GET("/getAllList",v2.GetAllInfoList)
		GroupV2.GET("/searchVideos",v2.SearchVideoHtml)
		GroupV2.POST("/searchVideos",v2.SearchVideoByName)
		GroupV2.GET("/videoPlayer",v2.PlayActionClick)
	}

}
