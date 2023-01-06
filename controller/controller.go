package controller

import (
	"account_book/models"
	"account_book/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 注册
func Register(c *gin.Context) {
	//从请求中拿出数据
	var user models.User
	c.ShouldBind(&user)
	//存入数据库
	if err := models.Register(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "创建成功"})

}

// 登录
func Login(c *gin.Context) {
	var user models.User
	c.ShouldBind(&user)
	users, err := models.Login(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": "用户名密码错误"})
		return
	}
	//使用jwt生成token
	token, _ := utils.ReleaseToken(user.Username)
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "登录成功", "data": users, "token": token})
}

// 修改密码
func UpdatePassword(c *gin.Context) {
	var user models.User
	c.ShouldBind(&user)
	err := models.IsCheckUser(user.Username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": "用户不存在"})
		return
	}
	//如果用户名存在修改密码
	if err = models.UpdatePassword(user.Username, user.Password); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "修改成功"})
	}
	//fmt.Println(c.Get("user"))
}
