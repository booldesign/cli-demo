package mysql

import "gorm.io/gorm"

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/12/2 17:10
 * @Desc:
 */

func SetPageSize(pageSize int) {
	defaultPageSize = pageSize
}

func SetMaxPageSize(pageSize int) {
	maxPageSize = pageSize
}

// DBLimit sql limit
func DBLimit(pageNum int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageNum <= 0 {
			pageNum = defaultPageNum
		}
		switch {
		case pageSize <= 0:
			pageSize = defaultPageSize
		case pageSize > maxPageSize:
			pageSize = maxPageSize
		}
		return db.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	}
}
