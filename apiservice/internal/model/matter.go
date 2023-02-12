package model

import "time"

// 文件
type Matter struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Uuid      string    `json:"uuid" gorm:"unique"`
	PID       uint      `json:"pid" gorm:"index:idx_matter_pid"` //index should unique globally.
	UserID    uint      `json:"user_id" gorm:"index:idx_matter_user"`
	Username  string    `json:"username"`
	Dir       *bool     `json:"dir"`
	Name      string    `json:"name"`
	Sha256    string    `json:"sha256"`
	Size      uint      `json:"size" gorm:"default:0"`
	Privacy   *bool     `json:"privacy"`
	Path      string    `json:"path"`
	Ext       string    `json:"ext"`
	Times     uint      `json:"times" gorm:"default:0"`
	Parent    *Matter   `json:"parent" gorm:"-"`
	Children  []*Matter `json:"-" gorm:"-"`
	VisitTime time.Time `json:"visit_time"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
