package model

import (
	"time"
)

// TaskLog 任务日志
type TaskLog struct {
	ID        int64     `gorm:"primaryKey;unique;column:id;type:bigint(20);not null"`
	TaskID    int64     `gorm:"column:task_id;type:bigint(20);not null"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
}

// TableName get sql table name.获取数据库表名
func (m *TaskLog) TableName() string {
	return "task_log"
}
