package models

import "time"

type BaseModel struct {
	CreatedAt time.Time  `gorm:"type: datetime" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"type: datetime" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"index; type: datetime" json:"-"`
}

// 每个表的基本字段，时间相关的信息记录，被继承
