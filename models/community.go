package models

import "time"

type Community struct {
	CommunityID      int       `gorm:"column:community_id" json:"Cid" binding:"required"`
	CommunityName    string    `gorm:"column:community_name" json:"Cname"`
	Introduction     string    `gorm:"column:introduction" json:"intro" binding:"required"`
	CreateTime       time.Time `gorm:"column:create_time" json:"createT"`
	UpdateTime       time.Time `gorm:"column:update_time" json:"updateT"`
	CommunityTitleID int64     `json:"commTId" gorm:"column:community_title_id"`
}

type CommunityRes struct {
	CommunityID      int    `gorm:"column:community_id" json:"Cid" binding:"required"`
	CommunityName    string `gorm:"column:community_name" json:"Cname"`
	Introduction     string `gorm:"column:introduction" json:"intro" binding:"required"`
	CommunityTitleID int64  `json:"commTId" gorm:"column:community_title_id"`
}
type CommunityLike struct {
	UserID           int64 `json:"user_id" gorm:"column:user_id" `
	CommunityTitleID int64 `json:"commTId" gorm:"column:community_title_id"`
	LikeTo           int   `json:"LikeTo" gorm:"column:liketo" binding:"required"`
}
