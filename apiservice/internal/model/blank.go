package model

import (
	"database/sql/driver"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Tags []string

func (t *Tags) Scan(val interface{}) error {
	s := val.(string)
	if len(s) <= 2 {
		return nil
	}
	ss := strings.Split(string(s)[1:len(s)-1], ",")
	*t = ss
	return nil
}

func (t Tags) Value() (driver.Value, error) {
	return "{" + strings.Join(t, ",") + "}", nil
}

type Blank struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	Type      string         `json:"type"`
	Title     string         `json:"title"`
	Content   string         `json:"content"`
	Tags      Tags           `json:"tags" gorm:"type:text[];column:tags"`
	Matters   []Matter       `json:"matters" gorm:"many2many:blank_matters;"`
	UserID    uint           `json:"user_id"`
	User      User           `json:"user" gorm:"-"`
	UpdatedAt time.Time      `json:"updated_at"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type BlankMatter struct {
	ID       uint `json:"id" gorm:"primary_key"`
	MatterID uint `json:"matter_id" gorm:"uniqueIndex:idx_blank_matter"`
	BlankID  uint `json:"blank_id" gorm:"uniqueIndex:idx_blank_matter"`
}
