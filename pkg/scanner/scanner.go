package scanner

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sync"
)

type Credential struct {
	FilePath string
	Line     string
	Pattern  string
}

func ScanDirectory(dir string, cfg *Config) ([]Credential, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var credentials []Credential
	files, err := getFiles(dir, cfg.FileTypes)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			fileCreds, err := ScanFile(file, cfg.Patterns)
			if err == nil {
				mu.Lock()
				credentials = append(credentials, fileCreds...)
				mu.Unlock()
			}
		}(file)
	}

	wg.Wait()
	return credentials, nil
}

func getFiles(dir string, fileTypes []string) ([]string, error) {
	var files []string
	return files, filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && (info.Name() == "node_modules") {
			return filepath.SkipDir
		}
		for _, ft := range fileTypes {
			if filepath.Ext(path) == ft {
				files = append(files, path)
				break
			}
		}
		return nil
	})
}

func ScanFile(file string, patterns []string) ([]Credential, error) {
	var credentials []Credential
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	lines := string(content)
	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		for _, line := range regexp.MustCompile("\n").Split(lines, -1) {
			if re.FindString(line) != "" {
				credentials = append(credentials, Credential{
					FilePath: file,
					Line:     line,
					Pattern:  pattern,
				})
			}
		}
	}
	return credentials, nil
}
