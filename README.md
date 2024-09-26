# go-cred-scanner

A CLI tool for scanning project directories to detect exposed credentials and sensitive information such as hardcoded API keys, usernames, passwords, and more. The scanner uses regex patterns defined in a configuration file to identify potential security risks in source code files.

## Table of Contents

- [go-cred-scanner](#go-cred-scanner)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Configuration](#configuration)
  - [Usage](#usage)
    - [Build the Project](#build-the-project)
    - [Run the Scanner](#run-the-scanner)
    - [Example Command](#example-command)
  - [Output](#output)
  - [Logging](#logging)
  - [Exclusions](#exclusions)
  - [File Type Filtering](#file-type-filtering)
  - [License](#license)

## Features

- **Customizable Patterns**: Define your own regex patterns for detecting sensitive information in the `config.json` file.
- **File Type Filtering**: Scan specific programming languages by setting file extensions in the configuration file.
- **Exclusions**: Automatically excludes common directories like `node_modules`.
- **Output to File**: Save the scan results in a specified output file.
- **Logging**: All activities are logged for review and debugging.

## Prerequisites

1. **Go Installation**: Ensure you have Go installed on your machine. Download it from [golang.org](https://golang.org/dl/).
2. **Git Installation**: If you don't have Git installed, download it from [git-scm.com](https://git-scm.com/).

## Installation

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/your-username/credentials-scanner.git
   cd credentials-scanner
   ```

2. **Create Configuration File**: Create the `config.json` file in the project root directory. Here's a sample configuration:

   ```json
   {
     "patterns": {
       "apiKey": "([a-zA-Z0-9]{32}|[A-Za-z0-9_]{36})",
       "usernamePassword": "(?i)(username|user|password|pass|secret)\\s*[:=]\\s*([\"']?)([^\"'\\s]+)\\2",
       "awsCredentials": "(?i)(aws_access_key_id|aws_secret_access_key)\\s*[:=]\\s*([\"']?)([^\"'\\s]+)\\2",
       "token": "([a-zA-Z0-9]{40})"
     },
     "fileExtensions": [
       ".go",
       ".js",
       ".ts",
       ".py",
       ".env"
     ]
   }
   ```

## Configuration

The `config.json` file contains patterns for scanning sensitive information and the file types to scan. Customize this file according to your needs.

## Usage

### Build the Project

```bash
go build -o scanner ./cmd/main.go
```

### Run the Scanner

```bash
./scanner -dir path/to/scan -config config.json -output results.json
```

Replace `path/to/scan` with the actual directory you want to scan.

### Example Command

To scan the current directory and save results to `results.json`, run:

```bash
./scanner -dir . -config config.json -output results.json
```

## Output

The results will be saved in the specified output file (e.g., `results.json`). Open this file to review the detected sensitive information.

## Logging

The scanner logs its activity to `logs/app.log`. You can check this file for detailed logs of the scanning process.

## Exclusions

The scanner automatically excludes directories like `node_modules` and similar, ensuring that your scans are focused on relevant code.

## File Type Filtering

You can modify the `fileExtensions` in `config.json` to scan for specific file types relevant to your projects.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
