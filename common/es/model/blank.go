package model

import (
	"time"
)

type Blank struct {
	ID        uint        `json:"id"`
	Type      string      `json:"type"`
	Title     string      `json:"title"`
	Content   string      `json:"content"`
	Tags      []string    `json:"tags"`
	Matters   []*Resource `json:"matters"`
	MatterIDs []uint      `json:"matter_ids"`
	UserID    uint        `json:"user_id"`
	UpdatedAt time.Time   `json:"update_at"`
	CreatedAt time.Time   `json:"create_at"`
	DeletedAt time.Time   `json:"deleted_at"`
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
