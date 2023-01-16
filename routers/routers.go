package routers

import (
	"account_book/controller"
	"account_book/middleware"

	"github.com/gin-gonic/gin"
)

func SetRouters() *gin.Engine {

	//初始化路由
	r := gin.Default()
	r.Use(middleware.Cors()) //开启中间件 允许使用跨域请求
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcomeGo",
		})
	})

	//user
	userGrop := r.Group("user")

	{
		userGrop.POST("/register", controller.Register)
		userGrop.POST("/login", controller.Login)
		userGrop.PUT("/password", middleware.AuthMiddleware(), controller.UpdatePassword)
	}

	//type
	typeGrop := r.Group("group")

	{
		typeGrop.POST("/create", middleware.AuthMiddleware(), controller.CreateGroup)
		typeGrop.GET("/list", middleware.AuthMiddleware(), controller.GetAllGroup)
	}

	//bill
	billGrop := r.Group("bill")
	{
		billGrop.POST("/create", middleware.AuthMiddleware(), controller.CreateBill)
		billGrop.GET("/list", middleware.AuthMiddleware(), controller.GetBillList)
	}
	return r
}
