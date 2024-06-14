package handlers

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"github.com/dchest/captcha"
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
	var captchaResponse struct {
		CaptchaID     string `form:"captcha_id" binding:"required"`
		CaptchaAnswer string `form:"captcha_answer" binding:"required"`
	}

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("User data: %+v\n", user)

	if err := c.ShouldBind(&captchaResponse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Captcha data: %+v\n", captchaResponse)
	log.Printf("Captcha data: %+v\n", captchaResponse.CaptchaID)
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

	c.JSON(http.StatusOK, gin.H{"token": token, "userId": dbUser.ID})
}

func ReplacePwd(c *gin.Context)  {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("User data: %+v\n", user)  // 调试绑定数据

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

	if err := database.DB.Save(&dbUser).Error; err !=nil {
		log.Print("updating pwd error:%v\n",err)
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK,gin.H{"message":"pwd reset successfull!"})

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
