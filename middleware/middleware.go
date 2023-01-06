package middleware

import (
	"account_book/dao"
	"account_book/models"
	"account_book/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取authorization header
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  "获取token失败",
			})
			c.Abort()
			return
		}

		// 获取真正的token字符串
		tokenString = tokenString[7:]
		token, claims, err := utils.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": "token错误",
			})
			c.Abort()
			return
		}
		username := claims.Username
		db := dao.DB
		var user models.User
		db.Model(&models.User{}).Where("username = ?", username).First(&user)

		// 如果没有找到
		if user.ID == 0 {
			c.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": "token不存在",
			})
			c.Abort()
			return
		}
		// 如果存在
		c.Set("user", user)
		c.Next()

	}
}
