package inittask

import (
	"nagato/searchservice/internal/config"
	"nagato/searchservice/internal/db/data"
)

func Init() {
	config.Init("./application.yml")

	data.GetDataCenter()
}
