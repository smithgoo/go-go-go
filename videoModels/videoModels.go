package videoModels

import "gorm.io/gorm"

type VideoInfo struct {
	gorm.Model
	Title  string  `json:"title"`
	Cover string `json:"cover"`
	Content string `json:"content"`
	Address string `json:"address"`
	OtherName string `json:"otherName"`
	VideoDirector string `json:"videoDirector"`
	VideoMaincharacter  string  `json:"videoMaincharacter"`
	VideoType  string  `json:"videoType"`
	VideoArea  string  `json:"videoArea"`
	VideoLanguage  string  `json:"videoLanguage"`
	VideoReleaseTime  string  `json:"videoReleaseTime"`
	VideoUpdate  string  `json:"videoUpdate"`
}

func (VideoInfo) TableName() string {
	return "videos_videoinfo"
}
