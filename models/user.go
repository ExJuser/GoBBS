package models

type User struct {
	UserID   int64  `gorm:"index:idx_user_id;unique;not null" json:"user_id,string"`
	Username string `gorm:"index:idx_username;unique;type:varchar(64);not null" json:"username"`
	Password string `gorm:"type:varchar(64);not null" json:"-"`
}
