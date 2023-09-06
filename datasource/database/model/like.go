package model

import "gorm.io/gorm"

type Like struct {
	Id        string         `gorm:"primaryKey"`
	VideoID   int64          `gorm:"index"`
	Who       int64          `gorm:"index"`
	IsLike    bool           `gorm:"not null"`
	UpdatedAt int64          `gorm:"autoUpdateTime"`
	CreatedAt int64          `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
