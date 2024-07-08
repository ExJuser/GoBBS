package models

import (
	"time"
)

type Community struct {
	ID   int64  `json:"community_id" gorm:"index:idx_community_id;column:community_id;not null;unique;"`
	Name string `json:"community_name" gorm:"index:idx_community_name;column:community_name;type:varchar(128);not null;unique"`
}

func (Community) TableName() string {
	return "community"
}

type CommunityDetail struct {
	ID           int64     `json:"community_id" gorm:"index:idx_community_id;column:community_id;not null;unique;"`
	Name         string    `json:"community_name" gorm:"index:idx_community_name;column:community_name;type:varchar(128);not null;unique"`
	Introduction string    `json:"introduction" gorm:"type:varchar(256);not null"`
	CreateTime   time.Time `json:"create_time"`
}

func (CommunityDetail) TableName() string {
	return "community"
}
