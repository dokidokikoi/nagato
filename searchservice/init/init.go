package inittask

import "nagato/searchservice/internal/db/data"

func Init() {
	data.GetDataCenter()
}
