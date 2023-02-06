package model

import "time"

// 文件
type Matter struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Uuid      string    `json:"uuid" gorm:"type:char(36);unique"`
	Puuid     string    `json:"puuid" gorm:"type:char(36);index:idx_matter_puuid"` //index should unique globally.
	UserUuid  string    `json:"user_uuid" gorm:"type:char(36);index:idx_matter_uu"`
	Username  string    `json:"username" gorm:"type:varchar(45) not null"`
	Dir       bool      `json:"dir" gorm:"type:tinyint(1) not null;default:0"`
	Name      string    `json:"name" gorm:"type:varchar(255) not null"`
	Sha256    string    `json:"sha256" gorm:"type:varchar(45)"`
	Size      int64     `json:"size" gorm:"type:bigint(20) not null;default:0"`
	Privacy   bool      `json:"privacy" gorm:"type:tinyint(1) not null;default:0"`
	Path      string    `json:"path" gorm:"type:varchar(1024)"`
	Ext       string    `json:"ext"`
	Times     int64     `json:"times" gorm:"type:bigint(20) not null;default:0"`
	Parent    *Matter   `json:"parent" gorm:"-"`
	Children  []*Matter `json:"-" gorm:"-"`
	VisitTime time.Time `json:"visit_time"`
	UpdateAt  time.Time `json:"update_at"`
	CreateAt  time.Time `json:"create_at"`
	DeleteAt  time.Time `json:"delete_at"`
}
