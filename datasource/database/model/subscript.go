package model

import "gorm.io/gorm"

type Sub struct {
	Id        string         `gorm:"primaryKey"`
	Who       int64          `gorm:"index"`
	IsSub     bool           `gorm:"default:false"`
	ToUserID  int64          `gorm:"index"`
	UpdatedAt int64          `gorm:"autoUpdateTime"`
	CreatedAt int64          `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
