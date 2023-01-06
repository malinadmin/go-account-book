package models

import (
	"account_book/dao"
	"time"
)

// 定义分类表模型
type Group struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"unique;not null" json:"name" form:"name"`         //分类的中文名称
	Type      int64  `gorm:"not null;check:type < 3" json:"type" form:"type"` //分类type1为支出 2为收入
	Userid    string
}

// 创建分类
func CreateGroup(group *Group) (err error) {
	err = dao.DB.Create(&group).Error
	return
}
