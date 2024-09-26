package test

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	config, err := scanner.LoadConfig("config.json")
	if err != nil {
		t.Fatalf("failed to load config: %s", err)
	}
	if len(config.Patterns) == 0 {
		t.Fatal("expected patterns to be loaded")
	}
}
