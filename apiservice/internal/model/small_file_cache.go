package model

import "time"

// 图片、文档等小文件可以多存一个完整的副本供预览和下载
// 做冷热数据处理，将长时间不访问的冷数据副本删除

var (
	SizeLimt = 10 << 20
)

type SmallFileCache struct {
	ID         uint      `json:"-" gorm:"primary_key"`
	MatterID   uint      `json:"-" gorm:"uniqueIndex"`
	Name       string    `json:"-" gorm:"index:idx_small_file_cache_name"`
	LastVisit  time.Time `json:"-"`
	VisitTimes uint      `json:"-"`
}
