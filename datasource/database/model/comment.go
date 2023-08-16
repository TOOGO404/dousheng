package model

type Comment struct {
	ID         int64  `gorm:"primaryKey;autoIncrement"`
	VideoID    int64  `gorm:"index "`
	Who        int64  `gorm:" "`
	Content    string `gorm:"default:''"`
	CreateDate string `gorm:" autoCreateTime"`
	User       User   `gorm:"foreignKey:Who"`
}
