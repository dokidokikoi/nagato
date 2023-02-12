package model

import "time"

type Blank struct {
	ID        uint      `json:"id"`
	Type      string    `json:"type"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Tags      []string  `json:"tags" gorm:"type:text[]"`
	MatterID  uint      `json:"matter_id"`
	Matters   []Matter  `json:"matters" gorm:"-"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type BlankMatter struct {
	ID       uint `json:"id" gorm:"primary_key"`
	MatterID uint `json:"matter_id" gorm:"uniqueIndex:idx_blank_matter"`
	BlankID  uint `json:"blank_id" gorm:"uniqueIndex:idx_blank_matter"`
}
