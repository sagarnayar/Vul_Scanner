VulScanner

VulScanner is a Go-based service that scans a GitHub repository for JSON files containing vulnerability data, stores them in an SQLite database, and allows querying based on severity.

Features

Scan API (POST /scan): Fetches and processes JSON vulnerability reports from a GitHub repository.

Query API (POST /query): Retrieves stored vulnerabilities filtered by severity.

Concurrency: Processes multiple files in parallel.

Error Handling: Retries GitHub API calls (up to 2 times) upon failure.

Docker Support: Runs as a single-container service.

Testing: Achieves 70%+ test coverage.

Installation & Setup

Prerequisites

Go (>=1.19)

SQLite3

Docker (optional, for containerized deployment)

Clone the Repository

git clone https://github.com/yourusername/vulscanner.git
cd vulscanner

Install Dependencies

go mod tidy

Run the Service

go run main.go

Run with Docker

docker build -t vulscanner .
docker run -p 8080:8080 vulscanner

API Usage

Scan API (POST /scan)

Request:

{
  "repo": "velancio/vulnerability_scans",
  "files": ["vulnscan1011.json"]
}

Response:

HTTP 200 OK
Scan completed

Query API (POST /query)

Request:

{
  "filters": {
    "severity": "HIGH"
  }
}

Response:

[
  {
    "id": "CVE-2024-1234",
    "severity": "HIGH",
    "cvss": 8.5,
    "status": "fixed",
    "package_name": "openssl",
    "current_version": "1.1.1t-r0",
    "fixed_version": "1.1.1u-r0",
    "description": "Buffer overflow vulnerability in OpenSSL",
    "published_date": "2024-01-15T00:00:00Z",
    "link": "https://nvd.nist.gov/vuln/detail/CVE-2024-1234",
    "risk_factors": ["Remote Code Execution", "High CVSS Score", "Public Exploit Available"]
  }
]

Testing

Run Unit Tests

go test ./... -v

Check Test Coverage

go test ./... -cover

Expected output:

vulscanner/database     coverage: 76.9%
vulscanner/handlers     coverage: 71.0%

Contribution & Deployment

Fork the repository and create a new branch.

Submit a pull request with clear commit messages.

Deploy by running docker build and docker run as shown above.
