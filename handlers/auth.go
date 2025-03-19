package handlers

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/dchest/captcha"
	"go-go-go/database"
	"go-go-go/models"
	"go-go-go/utils"
	"gorm.io/gorm"
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

func ShowResetPwdPage(c *gin.Context) {
	c.HTML(http.StatusOK, "resetPwd.html", nil)
}

func hashPassword(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}

func Register(c *gin.Context) {
	var user models.User
	var captchaResponse struct {
		CaptchaID     string `form:"captcha_id" binding:"required"`
		CaptchaAnswer string `form:"captcha_answer" binding:"required"`
	}

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBind(&captchaResponse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 验证验证码
	if !captcha.VerifyString(captchaResponse.CaptchaID, captchaResponse.CaptchaAnswer) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid captcha"})
		return
	}

	// 对密码进行 MD5 哈希处理
	user.Password = hashPassword(user.Password)

	// 查找 "viewer" 角色
	var viewerRole models.Role
	if err := database.DB.Where("name = ?", models.ViewerRole).First(&viewerRole).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "viewer role not found"})
		return
	}

	// 分配 "viewer" 角色
	user.RoleID = viewerRole.ID

	log.Printf("User data: %+v\n", user) // 调试绑定数据

	if err := database.DB.Create(&user).Error; err != nil {
		log.Printf("Error inserting user: %v\n", err) // 调试数据库错误
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration successful"})
}

func Login(c *gin.Context) {
	var user models.User
	var captchaResponse struct {
		CaptchaID     string `form:"captcha_id" binding:"required"`
		CaptchaAnswer string `form:"captcha_answer" binding:"required"`
	}

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Captcha data: %+v\n", captchaResponse)
	log.Printf("Captcha data: %+v\n", captchaResponse.CaptchaID)
	log.Printf("User data: %+v\n", user)
	log.Printf("Captcha data: %+v\n", captchaResponse)
	log.Printf("Captcha data: %+v\n", captchaResponse.CaptchaID)

	if err := c.ShouldBind(&captchaResponse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 验证验证码
	// 检查 CaptchaID 和 CaptchaAnswer 是否为空
	if captchaResponse.CaptchaID == "" || captchaResponse.CaptchaAnswer == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "captcha_id and captcha_answer are required"})
		return
	}

	// 验证验证码
	if !captcha.VerifyString(captchaResponse.CaptchaID, captchaResponse.CaptchaAnswer) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid captcha"})
		return
	}

	// 获取表单中输入的邮箱或手机号
	emailOrPhone := user.Email

	var dbUser models.User
	// 使用邮箱或手机号进行查询
	if err := database.DB.Where("email = ? OR phone = ?", emailOrPhone, emailOrPhone).First(&dbUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// 对输入的密码进行 MD5 哈希处理
	hashedPassword := hashPassword(user.Password)
	//log.Printf("hashedPassword: %+v\n", hashedPassword)  // 调试绑定数据
	//log.Printf("shashedPassword: %+v\n", dbUser.Password)  // 调试绑定数据
	if dbUser.Password != hashedPassword {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := utils.GenerateJWT(dbUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "user": dbUser})
}

func ReplacePwd(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("User data: %+v\n", user) // 调试绑定数据

	phone := user.Phone
	email := user.Email
	pwd := user.Password

	var dbUser models.User
	// 使用邮箱或手机号进行查询
	if err := database.DB.Where("email = ? AND phone = ?", email, phone).First(&dbUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	hashedPassword := hashPassword(pwd)
	dbUser.Password = hashedPassword

	if err := database.DB.Save(&dbUser).Error; err != nil {
		log.Print("updating pwd error:%v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "pwd reset successfull!"})

}

// GenerateCaptcha 生成验证码并返回验证码 ID 和验证码图像
func GenerateCaptcha(c *gin.Context) {
	captchaID := captcha.New()
	var content bytes.Buffer
	if err := captcha.WriteImage(&content, captchaID, 240, 80); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate captcha image"})
		return
	}
	captchaImage := base64.StdEncoding.EncodeToString(content.Bytes())

	c.JSON(http.StatusOK, gin.H{
		"captcha_id":    captchaID,
		"captcha_image": "data:image/png;base64," + captchaImage,
	})
}

// 新增方法
type UpdateUserRoleRequest struct {
	UserID  uint `json:"user_id" binding:"required"`  // 从请求中获取的 user_id
	RoleID  uint `json:"role_id" binding:"required"`  // 从请求中获取的 role_id
	IsAdmin uint `json:"is_admin" binding:"required"` // 从请求中获取的 role_id
}

// 获取所有用户信息
func GetAllUsers(c *gin.Context) {
	var users []models.User

	if err := database.DB.Preload("Role").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func UpdateUserRole(c *gin.Context) {
	var request UpdateUserRoleRequest
	var user models.User

	// 解析请求体，获取 user_id 和 role_id
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// 查找用户
	if err := database.DB.First(&user, request.UserID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		}
		return
	}

	// 打印用户的RoleID，确保用户已经从数据库获取
	println("&&&&&&&&&&&&&&&&&")
	// 打印整个用户结构体
	fmt.Printf("User: %+v\n", user)
	println("&&&&&&&&&&&&&&&&&")

	// 如果不是管理员，不允许修改权限
	if request.IsAdmin != 1 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		return
	}

	// 更新用户的角色ID
	user.RoleID = request.RoleID

	// 保存用户信息
	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User role updated successfully!"})
}
