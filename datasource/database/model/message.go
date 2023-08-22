package model

import "time"

type Message struct {
	Id         int64     `gorm:"primaryKey;autoIncrement"`
	FromUser   int64     `gorm:" "`
	ToUser     int64     `gorm:" "`
	Content    string    `gorm:"default:''"`
	CreateTime time.Time `gorm:" autoCreateTime"`
	User       User      `gorm:" foreignKey:ToUser"`
	User2      User      `gorm:" foreignKey:FromUser"`
}

//ID         int64     `gorm:"primaryKey;autoIncrement"`
//VideoID    int64     `gorm:"index "`
//Who        int64     `gorm:" "`
//Content    string    `gorm:"default:''"`
//CreateDate time.Time `gorm:" autoCreateTime"`
//User       User      `gorm:"foreignKey:Who"`
