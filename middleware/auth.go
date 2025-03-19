package middleware

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go-go-go/database"
	"go-go-go/models"
	"go-go-go/utils" // 导入你的 utils 包，确保路径正确
	"log"
	"net/http"
	"time"
)

// Redis 客户端
var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis 服务器地址
	})
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 Authorization 头部的值
		tokenStr := c.GetHeader("Authorization")
		// 如果 tokenStr 为空，返回 401 Unauthorized
		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			c.Abort()
			return
		}

		// 检查是否包含 "Bearer " 前缀
		if len(tokenStr) > 7 && tokenStr[:7] == "Bearer " {
			tokenStr = tokenStr[7:] // 去除前缀
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
			c.Abort()
			return
		}

		// 验证 token
		claims, err := utils.ValidateJWT(tokenStr)
		if err != nil {
			// 如果 token 过期，返回相应的错误信息
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorExpired != 0 {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
					c.Abort()
					return
				}
			}
			println("*********")
			println(claims)
			println("*********")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// 设置用户 ID 到上下文中
		c.Set("userID", claims.UserID)
		c.Next()
	}
}

func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		// 检查 IP 是否已注册
		val, err := redisClient.Get(context.Background(), "register:"+ip).Result()
		log.Printf("key: %v", val)
		if err == redis.Nil {
			// IP 不存在，允许继续
			redisClient.Set(context.Background(), "register:"+ip, "1", 24*time.Hour)
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Redis error"})
			c.Abort()
			return
		} else {
			// IP 已存在，限制注册
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func PreventDuplicateRegistration() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		var existingUser models.User
		if err := database.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
			c.Abort()
			return
		}

		c.Next()
	}
}
