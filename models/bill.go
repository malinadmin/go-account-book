package models

import (
	"account_book/dao"
	"time"
)

// 定义账单表模型
type Bill struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Amount    int64  `gorm:"not null" json:"amount" form:"amount"`
	TypeId    int64  `gorm:"not null" json:"type_id" form:"type_id"`
	TypeName  string `gorm:"not null" json:"type_name" form:"type_name"`
	GroupType int64  `gorm:"not null" json:"group_type" form:"group_type"` //分类type1为支出 2为收入
	Remark    string `json:"remark" form:"remark"`
	UserId    uint   `gorm:"not null" json:"user_id" form:"user_id"`
}

// 创建账单表
func CreateBill(bill *Bill) (err error) {
	if err = dao.DB.Create(&bill).Error; err != nil {
		return err
	}
	return
}

// 获取账单例表
func GetBillList(user_id uint, group_type string, pageSize int, offset int) (billList []*Bill, err error) {
	if err = dao.DB.Where("user_id = ? AND group_type = ?", user_id, group_type).Limit(pageSize).Offset(offset).Find(&billList).Error; err != nil {
		return nil, err
	}

	return
}
