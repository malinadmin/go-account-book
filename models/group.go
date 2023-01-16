package models

import (
	"account_book/dao"
	"time"
)

// 定义分类表模型
type Group struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"ureated_at"`
	Name      string    `gorm:"unique;not null" json:"name" form:"name"`         //分类的中文名称
	Type      int64     `gorm:"not null;check:type < 3" json:"type" form:"type"` //分类type1为支出 2为收入
	Userid    string
}

// 创建分类
func CreateGroup(group *Group) (err error) {
	err = dao.DB.Create(&group).Error
	return
}

// 获取所有分类
func GetAllGroup() (group *[]Group, err error) {
	if err = dao.DB.Find(&group).Error; err != nil {
		return nil, err
	}
	return
}
