package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"go-go-go/handlers"
	"go-go-go/middleware"
	scrapy1 "go-go-go/scrapy"
	v1 "go-go-go/v1"
	v2 "go-go-go/v2"
	"strings"
	"time"
)
import _ "net/url"
import _ "strconv"

func addCaseInsensitiveRoute(group *gin.RouterGroup, method, path string, handler gin.HandlerFunc) {
	// 注册原始路径
	group.Handle(method, path, handler)

	// 注册小写路径
	lowerPath := strings.ToLower(path)
	if lowerPath != path {
		group.Handle(method, lowerPath, handler)
	}
}

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

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, X-CSRF-Token")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	GroupV3 := r.Group("/user")
	{

		GroupV3.GET("/register", handlers.ShowRegisterPage)
		GroupV3.GET("/login", handlers.ShowLoginPage)

		//GroupV3.POST("/register", middleware.PreventDuplicateRegistration(), middleware.RateLimit(), handlers.Register)
		GroupV3.POST("/register", handlers.Register)
		GroupV3.POST("/login", handlers.Login)

		GroupV3.GET("/resetPwd", handlers.ShowResetPwdPage)
		GroupV3.POST("/resetPwd", handlers.ReplacePwd)

		GroupV3.GET("/captcha", handlers.GenerateCaptcha)

		GroupV3.GET("/home", middleware.AuthMiddleware(), handlers.Home)

		//获取用户列表
		GroupV3.GET("/getAllUser", middleware.AuthMiddleware(), handlers.GetAllUsers)

		// 编辑用户权限角色
		GroupV3.POST("/edit/role", middleware.AuthMiddleware(), handlers.UpdateUserRole)

	}

	GroupV1 := r.Group("/v1")
	{
		GroupV1.GET("/homePage", middleware.AuthMiddleware(), v1.ShowHomePage)
		GroupV1.GET("/product/add", v1.ShowAddProductPage)
		GroupV1.GET("/products", v1.ListProducts)
		GroupV1.POST("/product/add", v1.AddProduct)
		GroupV1.GET("/product/:id", v1.GetProduct)
		GroupV1.DELETE("/product/:id", v1.DeleteProduct)
		GroupV1.GET("/product/edit/:id", v1.ShowEditProductPage)
		GroupV1.POST("/product/edit/:id", v1.EditProduct)
		GroupV1.GET("/products/search", v1.SearchGetProductWithKw)
		GroupV1.POST("/products/search", v1.SearchProductWithKw)
	}

	GroupV2 := r.Group("/v2")
	{
		GroupV2.GET("/getAllList", v2.GetAllInfoList)
		GroupV2.GET("/getHomeAll", v2.GetHomeAllInfoList)
		GroupV2.GET("/searchVideos", v2.SearchVideoHtml)
		GroupV2.POST("/searchVideos", v2.SearchVideoByName)
		GroupV2.GET("/videoPlayer", v2.PlayActionClick)
	}

	GroupV4 := r.Group("/v4")
	{
		GroupV4.GET("/getScrapyInfo", scrapy1.FetchContent)
		//GroupV4.GET("/getVideoDetail",scrapy1.FetchCurrentVideoInfo)
		addCaseInsensitiveRoute(GroupV4, "GET", "/getVideoDetail", scrapy1.FetchCurrentVideoInfo)
	}

}
