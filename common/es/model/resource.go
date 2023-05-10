package model

import (
	"time"
)

type Resource struct {
	ID        uint        `json:"id"`
	UUID      string      `json:"uuid,omitempty"`
	PUUID     string      `json:"puuid,omitempty"`
	UserID    uint        `json:"user_id"`
	Dir       bool        `json:"dir"`
	Name      string      `json:"name,omitempty"`
	Sha256    string      `json:"sha256,omitempty"`
	Size      uint        `json:"size"`
	Privacy   bool        `json:"privacy"`
	Path      string      `json:"path,omitempty"`
	Ext       string      `json:"ext,omitempty"`
	Times     uint        `json:"times"`
	Children  []*Resource `json:"children,omitempty"`
	VisitTime time.Time   `json:"visit_time,omitempty"`
	UpdatedAt time.Time   `json:"update_at,omitempty"`
	CreatedAt time.Time   `json:"create_at,omitempty"`
	DeletedAt time.Time   `json:"deleted_at,omitempty"`
}

type ResourceReq struct {
	Dir          *bool     `json:"dir"`
	Text         string    `json:"text"`
	Sha256       string    `json:"sha256"`
	Size         uint      `json:"size"`
	Privacy      bool      `json:"privacy"`
	Path         string    `json:"path"`
	TimesGte     *uint     `json:"times_gte"`
	TimesLt      *uint     `json:"times_lt"`
	Ext          string    `json:"ext"`
	Nested       string    `json:"nested"`
	Highlight    []string  `json:"highlight"`
	UpdatedAtGte time.Time `json:"update_at_gte"`
	UpdatedAtLt  time.Time `json:"update_at_lt"`
	CreatedAtGte time.Time `json:"create_at_gte"`
	CreatedAtLt  time.Time `json:"create_at_lt"`
	Page         int       `form:"page"`
	PageSize     int       `form:"page_size"`
	Select       []string  `json:"select"` // es返回的sources包含的字段
}
