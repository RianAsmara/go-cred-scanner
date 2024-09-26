package scanner

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

func ExportResults(results []Credential, filename string, format string) error {
	switch format {
	case "csv":
		return exportToCSV(results, filename)
	case "json":
		return exportToJSON(results, filename)
	case "text":
		return exportToText(results, filename)
	default:
		return fmt.Errorf("unsupported format: %s", format)
	}
}

func exportToCSV(results []Credential, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	writer.Write([]string{"FilePath", "Line", "Pattern"})

	for _, result := range results {
		writer.Write([]string{result.FilePath, result.Line, result.Pattern})
	}

	return nil
}

func exportToJSON(results []Credential, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(results)
}

func exportToText(results []Credential, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, result := range results {
		if _, err := file.WriteString(result.FilePath + ": " + result.Line + " (Pattern: " + result.Pattern + ")\n"); err != nil {
			return err
		}
	}
	return nil
}
