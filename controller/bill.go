package controller

import (
	"account_book/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 新增账单
func CreateBill(c *gin.Context) {
	//从请求中拿出数据
	user, _ := c.MustGet("user").(models.User)

	var bill models.Bill
	bill.UserId = user.ID //从token中取出缓存数据
	c.ShouldBind(&bill)
	//存入数据库
	if err := models.CreateBill(&bill); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "创建成功"})

}

// 查询账单
func GetBillList(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", c.Query("page")))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", c.Query("page_size")))

	//从请求中拿出数据
	user, _ := c.MustGet("user").(models.User)
	group_type := c.Query("group_type") //类型1收入 2支出
	type_name := c.Query("type_name")
	total, list, err := models.GetBillList(user.ID, group_type, type_name, page, pageSize)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
		return
	}
	//获取总条数
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "请求成功", "data": map[string]interface{}{
		"data":     list,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	}})
}
