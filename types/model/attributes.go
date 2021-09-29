package model

import (
	"gintest/stores/mysql"
	"time"
)

// Attributes 属性
type Attributes struct {
	ID        int64     `json:"id" gorm:"primaryKey;unique;column:id;type:bigint(20);not null"`
	ParentID  int64     `json:"parent_id" gorm:"column:parent_id;type:bigint(20);default:0"`
	UserID    int64     `json:"user_id" gorm:"column:user_id;type:int(11);not null"`
	Name      string    `json:"name" gorm:"column:name;type:varchar(20);not null"`
	Sort      int       `json:"sort" gorm:"column:sort;type:int(11);default:0"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
}

// TableName get sql table name.获取数据库表名
func (m *Attributes) TableName() string {
	return "attributes"
}

func CreateAttribute(attribute Attributes) (ID int64, err error) {
	db, err := mysql.NewDb()
	if err != nil {
		return
	}
	db.Create(&attribute)
	if db.Error != nil {
		return 0, db.Error
	}

	return attribute.ID, nil
}

func UpdateAttribute(attributes Attributes) (err error) {
	db, err := mysql.NewDb()
	if err != nil {
		return
	}

	newAttribute := Attributes{ID: attributes.ID}
	db.Model(&newAttribute).Updates(attributes)

	if db.Error != nil {
		return db.Error
	}

	return
}
