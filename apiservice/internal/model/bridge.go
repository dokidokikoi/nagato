package model

/**
 * the link table for Share and Matter.
 */
type Bridge struct {
	ShareID  uint `json:"share_id" gorm:"uniqueIndex:idx_share_matter"`
	MatterID uint `json:"matter_id" gorm:"uniqueIndex:idx_share_matter"`
}
