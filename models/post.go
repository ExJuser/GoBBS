package models

import (
	"time"
)

type Post struct {
	ID          int64     `json:"id" gorm:"column:post_id;not null"`
	AuthorID    int64     `json:"authorID" gorm:"not null"`
	CommunityID int64     `json:"communityID" gorm:"not null" binding:"required"`
	Status      int32     `json:"status" gorm:"not null"`
	Title       string    `json:"title" gorm:"not null" binding:"required"`
	Content     string    `json:"content" gorm:"not null" binding:"required"`
	CreateTime  time.Time `json:"createTime"`
}

func (Post) TableName() string {
	return "post"
}
