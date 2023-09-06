package model

import "gorm.io/gorm"

type User struct {
	Id        int64          `gorm:"primaryKey;<-:create;autoIncrement"`
	Email     string         `gorm:"index;unique;size:128"`
	Pwd       string         `gorm:"not null"`
	Name      string         `gorm:"default:''"`
	Avatar    string         `gorm:"not null"`
	Backgroud string         `gorm:"not null"`
	Signature string         `gorm:"default:'这个人很懒，什么都没留下'"`
	WorkCount int64          `gorm:"default:0"`
	UpdatedAt int64          `gorm:"autoUpdateTime"`
	CreatedAt int64          `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
