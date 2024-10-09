package model

import "gorm.io/gorm"
import "gorm.io/datatypes"

// Bangumi 番剧
type Bangumi struct {
	gorm.Model
	SubscriptionID uint           `gorm:"column:subscription_id"`
	Subscription   *Subscription  `gorm:"foreignKey:SubscriptionID"` // 订阅
	OfficialTitle  string         `gorm:"column:official_title"`     // 名称
	FirstAirData   string         `gorm:"column:first_air_data"`     // 首播日期
	Season         int            `gorm:"column:season"`             // 季度
	GroupName      string         `gorm:"column:group_name"`         // 字幕组/发布者
	Poster         string         `gorm:"column:poster"`             // 海报
	Meta           datatypes.JSON `gorm:"column:meta"`               // 元数据
}

// BangumiEpisode 番剧集
type BangumiEpisode struct {
	gorm.Model
	SubscriptionFileID uint              `gorm:"column:subscription_file_id"`
	SubscriptionFile   *SubscriptionFile `gorm:"foreignKey:SubscriptionFileID"` // 订阅文件项
	Name               string            `gorm:"column:name"`                   // 名称
	Episode            int               `gorm:"column:episode"`                // 集数
	Poster             string            `gorm:"column:poster"`                 // 海报
	Overview           string            `gorm:"column:overview"`               // 简介
}
