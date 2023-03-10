package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             uint           `json:"id" gorm:"primarykey"`
	UUID           string         `json:"uuid" gorm:"uniqueIndex:idx_user_uuid"`
	Username       string         `json:"username" gorm:"unique"`
	Password       string         `json:"-"`
	Email          string         `json:"email" gorm:"unique"`
	Avatar         string         `json:"avatar"`
	LastIp         string         `json:"last_ip"`
	LastLogin      time.Time      `json:"last_login"`
	SizeLimit      int64          `json:"size_limit" gorm:"default:-1"`
	TotalSizeLimit int64          `json:"total_size_limit" gorm:"default:-1"`
	TotalSize      int64          `json:"total_size" gorm:"default:0"`
	Status         int            `json:"status"`
	UpdatedAt      time.Time      `json:"updated_at"`
	CreatedAt      time.Time      `json:"created_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at,omitempty"`
}
