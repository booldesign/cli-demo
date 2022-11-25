package usercenter

import "time"

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2022/11/25 11:01
 * @Desc:
 */

// LogModel 操作日志
type LogModel struct {
	ID         int64     `gorm:"primary_key;column:id"`
	Event      string    `gorm:"column:event"`       // 事件
	CreateTime time.Time `gorm:"column:create_time"` // 插入时间
	Data       string    `gorm:"column:data"`        // 内容
}

func (LogModel) TableName() string {
	return "log"
}
