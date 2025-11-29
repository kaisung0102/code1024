package model

type Community struct {
	ID            int64  `json:"id" gorm:"primaryKey"`
	CommunityID   int64  `json:"community_id" gorm:"uniqueIndex:idx_community_id"`
	CommunityName string `json:"community_name" gorm:"uniqueIndex:idx_community_name"`
	Introduction  string `json:"introduction"`
}
