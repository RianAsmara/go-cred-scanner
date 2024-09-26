// internal/usecase/scanner.go
package usecase

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/RianAsmara/go-cred-scanner/pkg/scanner"
)

// ScannerUseCase contains methods for scanning directories.
type ScannerUseCase struct{}

// Excluded directories that should not be scanned.
var excludedDirs = []string{"node_modules", ".git", "vendor"}

// ScanDirectory scans the given directory for sensitive information based on specified patterns and file types.
func (s *ScannerUseCase) ScanDirectory(dir string, patterns []string, fileTypes []string) ([]scanner.Credential, error) {
	var credentials []scanner.Credential

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip excluded directories
		if isExcludedDir(path) {
			return filepath.SkipDir
		}

		// Check if the file has a valid extension
		if !info.IsDir() {
			if !isValidFileType(path, fileTypes) {
				return nil
			}

			// Scan the file for credentials
			fileCreds, err := scanner.ScanFile(path, patterns)
			if err != nil {
				return err
			}
			credentials = append(credentials, fileCreds...)
		}
		return nil
	})

	return credentials, err
}

// isValidFileType checks if the file has a valid extension.
func isValidFileType(filePath string, fileTypes []string) bool {
	ext := strings.ToLower(filepath.Ext(filePath))
	for _, ft := range fileTypes {
		if ext == ft {
			return true
		}
	}
	return false
}

// isExcludedDir checks if the directory should be excluded from scanning.
func isExcludedDir(path string) bool {
	for _, excluded := range excludedDirs {
		if strings.Contains(path, excluded) {
			return true
		}
	}
	return false
}
