package config

import "testing"

func Test_NewModuleConfig(t *testing.T) {
	confPaths = "test"
	NewModuleConfig()
}
