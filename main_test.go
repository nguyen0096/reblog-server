package main

import (
	"reblog-server/utils/config"
	"testing"
)

func TestMain(m *testing.M) {
	config.InitConfig()
}
