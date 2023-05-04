package model

import (
	"time"
)

type Resource struct {
	ID        uint        `json:"id"`
	UUID      string      `json:"uuid"`
	PUUID     string      `json:"puuid"`
	UserID    uint        `json:"user_id"`
	Dir       bool        `json:"dir"`
	Name      string      `json:"name"`
	Sha256    string      `json:"sha256"`
	Size      uint        `json:"size"`
	Privacy   bool        `json:"privacy"`
	Path      string      `json:"path"`
	Ext       string      `json:"ext"`
	Times     uint        `json:"times"`
	Children  []*Resource `json:"children"`
	VisitTime time.Time   `json:"visit_time"`
	UpdatedAt time.Time   `json:"update_at"`
	CreatedAt time.Time   `json:"create_at"`
	DeletedAt time.Time   `json:"deleted_at"`
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
