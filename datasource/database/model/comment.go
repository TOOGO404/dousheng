package model

import "gorm.io/gorm"

type Comment struct {
	ID        int64 `gorm:"primaryKey;autoIncrement"`
	VideoID   int64 `gorm:"index "`
	Who       int64 `gorm:"not null"`
	SendTo    int64
	Content   string         `gorm:"default:''"`
	UpdatedAt int64          `gorm:"autoUpdateTime"` // 使用时间戳毫秒数填充更新时间
	CreatedAt int64          `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
