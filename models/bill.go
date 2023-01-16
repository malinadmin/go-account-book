package models

import (
	"account_book/dao"
	"account_book/utils"
	"time"
)

// 定义账单表模型
type Bill struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
	Amount    int64     `gorm:"not null" json:"amount" form:"amount"`
	TypeId    int64     `gorm:"not null" json:"type_id" form:"type_id"`
	TypeName  string    `gorm:"not null" json:"type_name" form:"type_name"`
	GroupType int64     `gorm:"not null" json:"group_type" form:"group_type"` //分类type1为支出 2为收入
	Remark    string    `json:"remark" form:"remark"`
	UserId    uint      `gorm:"not null" json:"user_id" form:"user_id"`
}

// 创建账单表
func CreateBill(bill *Bill) (err error) {
	if err = dao.DB.Create(&bill).Error; err != nil {
		return err
	}
	return
}

// 获取账单例表
func GetBillList(user_id uint, group_type string, type_name string, page int, page_size int) (total int64, billList []*Bill, err error) {
	Db := dao.DB         //这儿要重新赋值给一个变量不能使用全局变量不然会出问题
	if type_name != "" { //需要的搜索条件
		Db = Db.Where("type_name = ?", type_name)
	}
	err = Db.Model(&Bill{}).Where("user_id = ? AND group_type = ? ", user_id, group_type).Count(&total).Error
	if err != nil {
		return 0, nil, err
	}
	err = Db.Scopes(utils.Paginate(page, page_size)).Where("user_id = ? AND group_type = ?", user_id, group_type).Find(&billList).Error

	return total, billList, err
}
