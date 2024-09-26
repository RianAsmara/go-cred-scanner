package test

import (
	"testing"

	"github.com/RianAsmara/go-cred-scanner/pkg/scanner"
)

func TestLoadConfig(t *testing.T) {
	config, err := scanner.LoadConfig("../config.json")
	if err != nil {
		t.Fatalf("failed to load config: %s", err)
	}
	if len(config.Patterns) == 0 {
		t.Fatal("expected patterns to be loaded")
	}
}
