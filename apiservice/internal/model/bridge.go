package model

import "time"

/**
 * the link table for Share and Matter.
 */
type Bridge struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	Uuid       string    `json:"uuid" gorm:"type:char(36);unique"`
	ShareUuid  string    `json:"shareUuid" gorm:"type:char(36)"`
	MatterUuid string    `json:"matterUuid" gorm:"type:char(36)"`
	CreateAt   time.Time `json:"create_at"`
}
