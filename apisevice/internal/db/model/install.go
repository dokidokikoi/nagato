package model

import "time"

const (
	Downloading = 1
	Success     = 2
	Failed      = 3
)

type Install struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	MatterUuid string    `json:"matter_uuid" gorm:"type:char(36);uniqueIndex"`
	UserUuid   string    `json:"user_uuid" gorm:"type:char(36);uniqueIndex"`
	Status     int       `json:"status" gorm:"default:1"`
	CreateAt   time.Time `json:"create_at"`
}
