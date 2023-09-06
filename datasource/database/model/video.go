package model

import "gorm.io/gorm"

type Video struct {
	Id        int64          `gorm:"primaryKey;autoIncrement"`
	Author    int64          `gorm:"index"`
	Title     string         `gorm:"size:252"`
	PlayUrl   string         `gorm:"not null"`
	CoverUrl  string         `gorm:"not null"`
	Status    int8           `gorm:"default:0"`
	UpdatedAt int64          `gorm:"autoUpdateTime"` // 使用时间戳毫秒数填充更新时间
	CreatedAt int64          `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

const (
	VIDEO_TRANSCODING int8 = iota //视频上传后
	VIDEO_CHECKING
	VIDEO_CHECK_FAILED
	VIDEO_PUBLISHED
)

var VideoCode2Msg map[int8]string

func init() {
	VideoCode2Msg = map[int8]string{
		VIDEO_TRANSCODING:  "视频转码中",
		VIDEO_CHECKING:     "视频审核中",
		VIDEO_CHECK_FAILED: "视频审核失败",
		VIDEO_PUBLISHED:    "视频发布成功",
	}
}
