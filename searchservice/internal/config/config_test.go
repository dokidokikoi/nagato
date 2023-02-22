package config

import (
	"fmt"
	"testing"
)

func TestInitConfig(t *testing.T) {
	Init("../conf/application.yml")
	fmt.Printf("config: %+v", Config())
}
