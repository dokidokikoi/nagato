package model

import (
	"time"
)

// 文件
type Matter struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	UUID   string `json:"uuid" gorm:"uniqueIndex:idx_matter_uuid"`
	PID    uint   `json:"pid" gorm:"default:0;index:idx_matter_pid;"` //index should unique globally.
	UserID uint   `json:"user_id" gorm:"default:0;index:idx_matter_user"`
	// Dir       *bool     `json:"dir"`
	Name   string `json:"name"`
	Sha256 string `json:"sha256" gorm:"index:idx_matter_sha256;"`
	Size   uint   `json:"size" gorm:"default:0"`
	// Privacy   *bool     `json:"privacy"`
	Path  string `json:"path" gorm:"uniqueIndex:idx_path"`
	Ext   string `json:"ext"`
	Times uint   `json:"times" gorm:"default:0"`
	// Parent    *Matter   `json:"parent" gorm:"-"`
	// Children  []*Matter `json:"children" gorm:"-"`
	VisitTime time.Time `json:"visit_time"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
	// User      *User     `json:"user" gorm:"-"`
}
