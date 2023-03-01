package model

import "time"

type Blank struct {
	ID        uint      `json:"id"`
	Type      string    `json:"type"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Tags      []string  `json:"tags" gorm:"type:text[]"`
	MatterIDs []uint    `json:"matter_ids" gorm:"type:int8[]"`
	Matters   []Matter  `json:"matters" gorm:"-"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// TODO: 弃用
type BlankMatter struct {
	ID       uint `json:"id" gorm:"primary_key"`
	MatterID uint `json:"matter_id" gorm:"uniqueIndex:idx_blank_matter"`
	BlankID  uint `json:"blank_id" gorm:"uniqueIndex:idx_blank_matter"`
}
