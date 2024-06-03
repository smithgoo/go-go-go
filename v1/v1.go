package v1

import (
	"github.com/gin-gonic/gin"
	"go-go-go/database"
	"go-go-go/models"
	"log"
	"net/http"
	"strconv"
)

func ShowHomePage(c *gin.Context)  {
	c.HTML(http.StatusOK,"index.html",gin.H{
		"title":"Home Page",
	})
}

func ShowAddProductPage(c *gin.Context)  {
	c.HTML(http.StatusOK,"add.html",gin.H{
		"title":"Add Product",
	})
}

func ShowEditProductPage(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.HTML(http.StatusOK, "edit.html", product)
}

func  AddProduct(c *gin.Context)  {
	var product models.Product
	product.Name = c.PostForm("name")
	price,err := strconv.ParseFloat(c.PostForm("price"),64)
	if err !=nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid price"})
		return
	}
	product.Price = price

	database.DB.Create(&product)
	
	//var existingProduct models.Product
	//if err := database.DB.Where("name = ?",product.Name).First(&existingProduct).Error; err ==nil {
	//	if c.GetHeader("Content-Type") == "application/json" {
	//		c.JSON(http.StatusBadRequest, gin.H{"error":"product is same name"})
	//	} else {
	//		c.HTML(http.StatusFound, "index.html",gin.H{"error":"product is same name"})
	//	}
	//	return
	//}

	// Check the "Accept" header to determine response type
	acceptHeader := c.GetHeader("Accept")
	if acceptHeader == "application/json" || acceptHeader == "application/xml" {
		c.JSON(http.StatusOK, product)
	} else {
		c.Redirect(http.StatusFound, "/v1/products")
	}

}

func DeleteProduct(c *gin.Context){
	id,err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid product ID"})
		return
	}

	var product models.Product
	if err := database.DB.First(&product,id).Error;  err != nil {
		c.JSON(http.StatusNotFound,gin.H{"error":"Product not found"})
		return
	}

	if err := database.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":"delete fail"})
		return
	}

	c.JSON(http.StatusOK,gin.H{"success":true})
}

func EditProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	product.Name = c.PostForm("name")
	price, err := strconv.ParseFloat(c.PostForm("price"), 64)
	if err != nil {
		if c.GetHeader("Content-Type") == "application/json" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price"})
		} else {
			c.HTML(http.StatusBadRequest, "edit.html", gin.H{"error": "Invalid price", "ID": id, "Name": product.Name, "Price": product.Price})
		}
		return
	}
	product.Price = price

	database.DB.Save(&product)

	// Check the "Accept" header to determine response type
	acceptHeader := c.GetHeader("Accept")
	if acceptHeader == "application/json" || acceptHeader == "application/xml" {
		c.JSON(http.StatusOK, product)
	} else {
		c.Redirect(http.StatusFound, "/v1/products")
	}
}

func SearchGetProductWithKw(c *gin.Context){
	c.HTML(http.StatusOK,"search.html",gin.H{
		"title":"search Product",
	})
}

func SearchProductWithKw(c *gin.Context){
	println("search action ~~~~~~~~~~~~~~~~")
	var products []models.Product
	keyword := c.PostForm("keyword")

	log.Println("Search keyword:", keyword)
	if keyword != "" {
		query := "%" + keyword + "%"
		log.Println("Search query:", query)
		database.DB.Where("name LIKE ?", query).Find(&products)
	} else {
		database.DB.Find(&products)
	}

	c.HTML(http.StatusOK, "search.html", gin.H{
		"products": products,
	})
}



func GetProduct(c *gin.Context){
	var product models.Product
	id := c.Param("id")
	if err := database.DB.Order("created_at desc").First(&product,id).Error; err!=nil {
		c.JSON(http.StatusNotFound,gin.H{"error":"Product not found"})
		return
	}
	c.JSON(http.StatusOK,product)
}



func ListProducts(c *gin.Context)  {
	var products []models.Product
	database.DB.Find(&products)
	c.HTML(http.StatusOK,"list.html",gin.H{
		"products":products,
	})
}
