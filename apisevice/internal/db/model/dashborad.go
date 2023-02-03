package model

import "time"

/**
 * application's dashboard.
 */
type Dashboard struct {
	Uuid           string    `json:"uuid" gorm:"type:char(36);primary_key;unique"`
	InvokeNum      int64     `json:"invoke_num" gorm:"type:bigint(20) not null"`                 //api invoke num.
	TotalInvokeNum int64     `json:"total_invoke_num" gorm:"type:bigint(20) not null;default:0"` //total invoke num up to now.
	Uv             int64     `json:"uv" gorm:"type:bigint(20) not null;default:0"`               //today's uv
	TotalUv        int64     `json:"total_uv" gorm:"type:bigint(20) not null;default:0"`         //total uv
	MatterNum      int64     `json:"matter_num" gorm:"type:bigint(20) not null;default:0"`       //file's num
	TotalMatterNum int64     `json:"total_matter_num" gorm:"type:bigint(20) not null;default:0"` //file's total number
	FileSize       int64     `json:"file_size" gorm:"type:bigint(20) not null;default:0"`        //today's file size
	TotalFileSize  int64     `json:"total_file_size" gorm:"type:bigint(20) not null;default:0"`  //total file's size
	AvgCost        int64     `json:"avg_cost" gorm:"type:bigint(20) not null;default:0"`         //api time cost in ms
	Dt             string    `json:"dt" gorm:"type:varchar(45) not null;index:idx_dashboard_dt"` //date. index should unique globally.
	UpdateAt       time.Time `json:"update_at"`
	CreateAt       time.Time `json:"create_at"`
}

/**
 * ip
 */
type DashboardIpTimes struct {
	Ip    string `json:"ip"`
	Times int64  `json:"times"`
}
