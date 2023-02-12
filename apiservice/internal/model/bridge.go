package model

import "time"

/**
 * the link table for Share and Matter.
 */
type Bridge struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	ShareID   uint      `json:"share_id" gorm:"uniqueIndex:idx_share_matter"`
	MatterID  uint      `json:"matter_id" gorm:"uniqueIndex:idx_share_matter"`
	CreatedAt time.Time `json:"created_at"`
}
