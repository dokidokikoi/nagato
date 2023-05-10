package model

import (
	"time"
)

type Blank struct {
	ID        uint        `json:"id,omitempty"`
	Type      string      `json:"type,omitempty"`
	Title     string      `json:"title,omitempty"`
	Content   string      `json:"content,omitempty"`
	Tags      []string    `json:"tags,omitempty"`
	Matters   []*Resource `json:"matters,omitempty"`
	MatterIDs []uint      `json:"matter_ids,omitempty"`
	UserID    uint        `json:"user_id,omitempty"`
	UpdatedAt time.Time   `json:"update_at,omitempty"`
	CreatedAt time.Time   `json:"create_at,omitempty"`
	DeletedAt time.Time   `json:"deleted_at,omitempty"`
}

type BlankReq struct {
	Type         string    `json:"type"`
	Text         string    `json:"text"`
	Tags         []string  `json:"tags"`
	Highlight    []string  `json:"highlight"`
	Nested       string    `json:"nested"`
	MatterIDs    []uint    `json:"matter_ids"`
	UpdatedAtGte time.Time `json:"update_at_gte"`
	UpdatedAtLt  time.Time `json:"update_at_lt"`
	CreatedAtGte time.Time `json:"create_at_gte"`
	CreatedAtLt  time.Time `json:"create_at_lt"`
	Page         int       `form:"page"`
	PageSize     int       `form:"page_size"`
	Select       []string  `json:"select"`
}
