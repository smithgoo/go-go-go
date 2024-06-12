package handlers

import (
	"go-go-go/models"
	"go-go-go/utils"
	"go-go-go/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowRegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func ShowLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("User data: %+v\n", user)  // 调试绑定数据

	if err := database.DB.Create(&user).Error; err != nil {
		log.Printf("Error inserting user: %v\n", err)  // 调试数据库错误
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration successful"})
}

func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("User data: %+v\n", user)  // 调试绑定数据

	// 获取表单中输入的邮箱或手机号
	emailOrPhone := user.Email

	var dbUser models.User
	// 使用邮箱或手机号进行查询
	if err := database.DB.Where("email = ? OR phone = ?", emailOrPhone, emailOrPhone).First(&dbUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	if dbUser.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := utils.GenerateJWT(dbUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "userId": dbUser.ID})
}

