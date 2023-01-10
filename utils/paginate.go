package utils

import (
	"gorm.io/gorm"
)

type Pagination struct {
	Current  int `json:"page" form:"page"`         // 页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页大小
}

// 分页查询
func Paginate(current int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if current == 0 {
			current = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (current - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
