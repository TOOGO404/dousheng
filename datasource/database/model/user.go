package model

type User struct {
	Id        int64   `gorm:"primaryKey;<-:create;autoIncrement"`
	Email     string  `gorm:"index;unique;size:128"`
	Pwd       string  `gorm:"size:64"`
	Name      string  `gorm:"default:''"`
	Avatar    string  `gorm:"default:''"`
	Backgroud string  `gorm:"default:''"`
	Signature string  `gorm:"default:''"`
	Updated   int64   `gorm:"autoUpdateTime:milli"` // 使用时间戳毫秒数填充更新时间
	Created   int64   `gorm:"autoCreateTime"`
	Videos    []Video `gorm:"foreignKey:Author"`
}
