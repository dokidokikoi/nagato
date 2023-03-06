package model

type Tag struct {
	ID      uint   `json:"id" gorm:"primarykey"`
	TagName string `json:"tag_name" gorm:"uniqueIndex"`
	UserID  uint   `json:"user_id"`
}
