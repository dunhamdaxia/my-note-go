package model

import (
	"time"
)

// AwardsLog 奖励日志
type AwardsLog struct {
	ID           int64     `gorm:"primaryKey;unique;column:id;type:bigint(20);not null"`
	UserID       int64     `gorm:"index:user_id_index;column:user_id;type:bigint(20);not null"`
	Number       int       `gorm:"column:number;type:int(11);default:0"`
	Desc         string    `gorm:"column:desc;type:varchar(255);default:''"`
	TaskID       int64     `gorm:"column:task_id;type:bigint(20);not null"`
	AttributesID int64     `gorm:"column:attributes_id;type:bigint(20);not null"`
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
}

// TableName get sql table name.获取数据库表名
func (m *AwardsLog) TableName() string {
	return "awards_log"
}
