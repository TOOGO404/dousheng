package model

type Like struct {
	VideoID int64 `gorm:"index "`
	Who     int64 `gorm:"index "`
	IsLike  bool  `gorm:"default:false "`
	User    User  `gorm:" foreignKey:Who"`
	Video   Video `gorm:" foreignKey:VideoID"`
}
