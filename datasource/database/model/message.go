package model

import "gorm.io/gorm"

type Message struct {
	Id        int64
	ComID     string `gorm:"index"`
	FromUser  int64
	ToUser    int64
	Msg       string
	UpdatedAt int64          `gorm:"autoUpdateTime:milli"`
	CreatedAt int64          `gorm:"autoCreateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"index:milli"`
}
