package controller

import (
	"account_book/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 创建分类
func CreateGroup(c *gin.Context) {
	var group models.Group
	c.ShouldBind(&group)
	if err := models.CreateGroup(&group); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "创建成功"})
}

// 获取所有分类
func GetAllGroup(c *gin.Context) {
	list, err := models.GetAllGroup()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": list})
}
