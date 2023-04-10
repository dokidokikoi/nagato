package model

import "time"

type Share struct {
	ID             uint      `json:"id" gorm:"primary_key"`
	UUID           string    `json:"uuid"`
	UserID         uint      `json:"user_id"`
	User           *User     `json:"user,omitempty" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	DownloadTimes  uint      `json:"download_times" gorm:"default:0"`
	Code           string    `json:"code"`            // 提取码
	ExpireInfinity bool      `json:"expire_infinity"` // 是否永不过期
	ExpireTime     time.Time `json:"expireTime"`
	Matters        []*Matter `json:"matters" gorm:"many2many:share_matters;"`
	CreatedAt      time.Time `json:"created_at"`
}

type ShareMatter struct {
	ShareID  uint `json:"share_id"`
	MatterID uint `json:"matter_id"`
}

func (s ShareMatter) TableName() string {
	return "share_matters"
}
