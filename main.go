package main

import (
	"account_book/dao"
	"account_book/models"
	"account_book/routers"
)

func main() {

	//连接数据库
	err := dao.InitMql()

	if err != nil {
		panic(err)
	}
	//模型绑定
	dao.DB.AutoMigrate(&models.User{}, &models.Group{}, &models.Bill{})
	//启动路由
	r := routers.SetRouters()
	r.Run()
}
