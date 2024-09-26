package cli

import (
	"flag"
	"fmt"

	"github.com/RianAsmara/go-cred-scanner/pkg/scanner"
)

func Execute() {
	dirPtr := flag.String("dir", ".", "Directory to scan for sensitive information")
	outputPtr := flag.String("output", "results.csv", "Output file name for results")
	formatPtr := flag.String("format", "csv", "Output format: csv, json, text")
	configPtr := flag.String("config", "config.json", "Configuration file")

	flag.Parse()

	// Load configuration
	cfg, err := scanner.LoadConfig(*configPtr)
	if err != nil {
		fmt.Printf("Error loading config: %s\n", err)
		return
	}

	// Scan directory
	results, err := scanner.ScanDirectory(*dirPtr, cfg)
	if err != nil {
		fmt.Printf("Error scanning directory: %s\n", err)
		return
	}

	// Export results
	err = scanner.ExportResults(results, *outputPtr, *formatPtr)
	if err != nil {
		fmt.Printf("Error exporting results: %s\n", err)
		return
	}

	fmt.Println("Scanning complete. Results saved to", *outputPtr)
}
