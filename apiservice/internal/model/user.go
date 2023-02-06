package model

import "time"

type User struct {
	ID             uint      `json:"id" gorm:"primarykey"`
	Username       string    `json:"username" gorm:"type:varchar(45) not null;unique"`
	Password       string    `json:"-" gorm:"type:varchar(255)"`
	Email          string    `json:"email" gorm:"type:varchar(255) not null;unique"`
	Avatar         string    `json:"avatar" gorm:"type:varchar(255)"`
	LastIp         string    `json:"last_ip" gorm:"type:varchar(128)"`
	LastLogin      time.Time `json:"last_login"`
	SizeLimit      int64     `json:"size_limit" gorm:"type:bigint(20) not null;default:-1"`
	TotalSizeLimit int64     `json:"total_size_limit" gorm:"type:bigint(20) not null;default:-1"`
	TotalSize      int64     `json:"total_size" gorm:"type:bigint(20) not null;default:0"`
	Status         int       `json:"status"`
	UpdateAt       time.Time `json:"update_at"`
	CreateAt       time.Time `json:"create_at"`
}
