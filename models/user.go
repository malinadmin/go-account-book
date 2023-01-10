package models

import (
	"account_book/dao"

	"gorm.io/gorm"
)

// 定义模型
type User struct {
	gorm.Model        //使用此内置参数删除时会使用软删除方式
	Username   string `gorm:"unique;not null" json:"username" form:"username" ` //username为唯一值并且不能为空
	Password   int64  `gorm:"not null" json:"password" form:"password"`
	Sign       string `json:"sign" form:"sign"`
	Avatr      string `json:"avatr" form:"avatr"`
}

// 注册
func Register(user *User) (err error) {
	if err = dao.DB.Create(&user).Error; err != nil {
		return err
	}
	return
}

// 检查用户是否存在
func IsCheckUser(username string) (err error) {
	var user User
	if err = dao.DB.Where(map[string]interface{}{"username": username}).First(&user).Error; err != nil {
		return err
	}
	return
}

// 登录
func Login(username string, password int64) (user *User, err error) {
	user = new(User)
	if err = dao.DB.Where(map[string]interface{}{"username": username, "password": password}).First(user).Error; err != nil {
		return nil, err
	}
	return
}

// 修改密码
func UpdatePassword(username string, password int64) (err error) {
	var user User
	err = dao.DB.Model(&user).Where("username = ?", username).Update("password", password).Error
	return
}
