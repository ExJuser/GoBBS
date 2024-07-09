package models

import (
	"time"
)

type Post struct {
	ID          int64     `json:"id,string" gorm:"column:post_id;not null"`
	AuthorID    int64     `json:"author_id,string" gorm:"not null"`
	CommunityID int64     `json:"community_id,string" gorm:"not null" binding:"required"`
	Status      int32     `json:"status" gorm:"not null"`
	Title       string    `json:"title" gorm:"not null" binding:"required"`
	Content     string    `json:"content" gorm:"not null" binding:"required"`
	CreateTime  time.Time `json:"create_time"`
}

func (Post) TableName() string {
	return "post"
}

type APIPostDetail struct {
	//还需要再返回作者信息
	AuthorName string           `json:"author_name"`
	Community  *CommunityDetail //社区信息
	Post       *Post            //具体帖子信息
}
