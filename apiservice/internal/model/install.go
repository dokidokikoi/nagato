package model

import "time"

const (
	Downloading = 1
	Success     = 2
	Failed      = 3
)

type Install struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	MatterID  uint      `json:"matter_id" gorm:"uniqueIndex:idx_user_matter"`
	UserID    uint      `json:"user_id" gorm:"uniqueIndex:idx_user_matter"`
	Status    int       `json:"status" gorm:"default:1"`
	CreatedAt time.Time `json:"created_at"`
}
