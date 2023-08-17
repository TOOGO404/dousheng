package model

import "time"

type Comment struct {
	ID         int64     `gorm:"primaryKey;autoIncrement"`
	VideoID    int64     `gorm:"index "`
	Who        int64     `gorm:" "`
	Content    string    `gorm:"default:''"`
	CreateDate time.Time `gorm:" autoCreateTime"`
	User       User      `gorm:"foreignKey:Who"`
}
