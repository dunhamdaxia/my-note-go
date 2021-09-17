package model

import (
	"time"
)

// Attributes 属性
type Attributes struct {
	ID        int64     `gorm:"primaryKey;unique;column:id;type:bigint(20);not null"`
	ParentID  int64     `gorm:"column:parent_id;type:bigint(20);default:0"`
	UserID    int       `gorm:"column:user_id;type:int(11);not null"`
	Name      string    `gorm:"column:name;type:varchar(20);not null"`
	Sort      int       `gorm:"column:sort;type:int(11);default:0"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
}

// TableName get sql table name.获取数据库表名
func (m *Attributes) TableName() string {
	return "attributes"
}
