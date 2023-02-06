package model

import "time"

type Share struct {
	ID             uint      `json:"id" gorm:"primary_key"`
	Uuid           string    `json:"uuid" gorm:"type:char(36);uniqueIndex"`
	Name           string    `json:"name" gorm:"type:varchar(255)"`
	ShareType      string    `json:"share_type" gorm:"type:varchar(45)"`
	Username       string    `json:"username" gorm:"type:varchar(45)"`
	UserUuid       uint      `json:"user_id"`
	DownloadTimes  uint      `json:"download_times" gorm:"not null;default:0"`
	Code           string    `json:"code" gorm:"type:varchar(45) not null"`
	ExpireInfinity bool      `json:"expire_infinity" gorm:"type:tinyint(1) not null;default:0"`
	ExpireTime     time.Time `json:"expireTime"`
	DirMatter      *Matter   `json:"dir_matter" gorm:"-"`
	Matters        []*Matter `json:"matters" gorm:"-"`
	CreateAt       time.Time `json:"create_at"`
}
