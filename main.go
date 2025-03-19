package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-go-go/database"
	"go-go-go/models"
	models2 "go-go-go/models"
	"go-go-go/router"
	"strings"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization, Accept, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Request.URL.Path = strings.ToLower(c.Request.URL.Path)
		c.Next()
	})

	r.Use(CORSMiddleware())
	// 中间件，将所有路径转换为小写

	fmt.Println("Current Gin mode:", gin.Mode())
	database.InitDB("root", "Jbnb123456", "127.0.0.1:3306", "DBInfo")

	database.DB.AutoMigrate(&models.Product{}, &models2.VideoInfo{}, &models2.User{})

	initializeRoles()

	router.InitRouter(r)

	r.Run(":9000")

}

func initializeRoles() {
	roles := []models.RoleType{models.AdminRole, models.UserRole, models.ViewerRole}
	for _, roleName := range roles {
		var role models.Role
		if err := database.DB.Where("name = ?", roleName).First(&role).Error; err != nil {
			database.DB.Create(&models.Role{Name: roleName})
		}
	}
}


