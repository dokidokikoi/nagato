package model

import (
	"time"

	"gorm.io/gorm"
)

// 文件
type Matter struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	UUID      string         `json:"uuid" gorm:"uniqueIndex:idx_matter_uuid"`
	PUUID     string         `json:"puuid" gorm:"index:idx_matter_puuid;default:null"` //index should unique globally.
	UserID    uint           `json:"user_id" gorm:"default:0;index:idx_matter_user"`
	Dir       bool           `json:"dir"`
	Name      string         `json:"name"`
	Sha256    string         `json:"sha256" gorm:"index:idx_matter_sha256;"`
	Size      uint           `json:"size" gorm:"default:0"`
	Privacy   bool           `json:"privacy"`
	Path      string         `json:"path" gorm:"uniqueIndex:idx_matter_path"`
	Ext       string         `json:"ext"`
	Times     uint           `json:"times" gorm:"default:0"`
	Parent    *Matter        `json:"parent,omitempty" gorm:"-"`
	Children  []*Matter      `json:"children" gorm:"foreignKey:PUUID;references:UUID"`
	VisitTime time.Time      `json:"visit_time"`
	UpdatedAt time.Time      `json:"updated_at"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
	User      *User          `json:"user,omitempty" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
