package models

type Community struct {
	ID   int64  `json:"community_id" gorm:"index:idx_community_id;column:community_id;not null;unique;"`
	Name string `json:"community_name" gorm:"index:idx_community_name;column:community_name;type:varchar(128);not null;unique"`
}

func (Community) TableName() string {
	return "user"
}
