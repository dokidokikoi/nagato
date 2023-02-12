package model

import "time"

type Share struct {
	ID             uint      `json:"id" gorm:"primary_key"`
	Name           string    `json:"name"`
	ShareType      string    `json:"share_type"`
	Username       string    `json:"username"`
	UserID         uint      `json:"user_id"`
	DownloadTimes  uint      `json:"download_times" gorm:"default:0"`
	Code           string    `json:"code"`
	ExpireInfinity *bool     `json:"expire_infinity"`
	ExpireTime     time.Time `json:"expireTime"`
	Matters        []*Matter `json:"matters" gorm:"-"`
	CreatedAt      time.Time `json:"created_at"`
}
