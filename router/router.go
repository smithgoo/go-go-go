package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	v2 "go-go-go/v2"
	v1 "go-go-go/v1"
)
import _ "net/url"
import _ "strconv"

func InitRouter(r *gin.Engine) {

	r.LoadHTMLGlob("templates/*")

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
	}

}
