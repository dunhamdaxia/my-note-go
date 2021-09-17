package model

import (
	"gintest/stores/mysql"
	"time"
)

// Task 任务表
type Task struct {
	ID         int64     `json:"id" gorm:"primaryKey;unique;column:id;type:bigint(20);not null"`
	UserID     int64     `json:"user_id" gorm:"column:user_id;type:bigint(20);not null"`
	Name       string    `json:"name" gorm:"column:name;type:varchar(100);not null"`
	Type       int8      `json:"type" gorm:"column:type;type:tinyint(4);default:1"`
	AwardsInfo string    `json:"awards_info" gorm:"column:awards_info;type:text;not null"`
	PeriodType int8      `json:"period_type" gorm:"column:period_type;type:tinyint(4);not null;default:1"` // 周期类型 1不重复 2每天 3每周 4每月
	Desc       string    `json:"desc" gorm:"column:desc;type:text"`                                        // 任务描述
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	StartDate  time.Time `json:"start_date" gorm:"column:start_date;type:timestamp;default:CURRENT_TIMESTAMP"`
	EndDate    time.Time `json:"end_date" gorm:"column:end_date;type:timestamp;not null;default:0000-00-00 00:00:00"`
	Status     int8      `json:"status" gorm:"column:status;type:tinyint(4);default:1"` // 状态 1待完成 2已完成
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
}

// TableName get sql table name.获取数据库表名
func (m *Task) TableName() string {
	return "task"
}

func CreateTask(task Task) (ID int64, err error) {
	db, err := mysql.NewDb()
	if err != nil {
		return
	}
	db.Create(&task)
	if db.Error != nil {
		return 0, db.Error
	}

	return task.ID, nil
}
