package model

type Video struct {
	Id       int64  `gorm:"primaryKey;autoIncrement"`
	Author   int64  `gorm:"index"`
	Title    string `gorm:"size:252"`
	PlayUrl  string
	CoverUrl string
	Status   int8
	Updated  int64 `gorm:"autoUpdateTime:milli"` // 使用时间戳毫秒数填充更新时间
	Created  int64 `gorm:"autoCreateTime"`
	User     User  `gorm:"foreignKey:Author"`
}

const (
	VIDEO_TRANSCODING int8 = iota //视频上传后
	VIDEO_CHECKING
	VIDEO_CHECK_FAILED
	VIDEO_PUBLISHED
	VIDEO_CANCELED
)

var VideoCode2Msg map[int8]string

func init() {
	VideoCode2Msg = map[int8]string{
		VIDEO_TRANSCODING:  "视频转码中",
		VIDEO_CHECKING:     "视频审核中",
		VIDEO_CHECK_FAILED: "视频审核失败",
		VIDEO_PUBLISHED:    "视频发布成功",
		VIDEO_CANCELED:     "视频取消发布",
	}
}
