package model

import (
	"database/sql/driver"
	commonEsModel "nagato/common/es/model"
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
	Matters   []*Matter      `json:"matters" gorm:"many2many:blank_matters;"`
	UserID    uint           `json:"user_id"`
	User      *User          `json:"user,omitempty" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	UpdatedAt time.Time      `json:"updated_at"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (b Blank) ToEsStruct(matterIDs []uint) commonEsModel.Blank {
	return commonEsModel.Blank{
		ID:        b.ID,
		Type:      b.Type,
		Title:     b.Title,
		Content:   b.Content,
		Tags:      b.Tags,
		MatterIDs: matterIDs,
		UserID:    b.UserID,
		UpdatedAt: b.UpdatedAt,
		CreatedAt: b.CreatedAt,
		DeletedAt: b.DeletedAt.Time,
	}
}

type BlankMatter struct {
	MatterID uint `json:"matter_id" gorm:"uniqueIndex:idx_blank_matter"`
	BlankID  uint `json:"blank_id" gorm:"uniqueIndex:idx_blank_matter"`
}
