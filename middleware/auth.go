package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-go-go/utils" // 导入你的 utils 包，确保路径正确
	"github.com/dgrijalva/jwt-go"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			c.Abort()
			return
		}

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

			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}
}


