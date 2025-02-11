package database

import (
	"database/sql"
	"log"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"vulscanner/models"
)

var DB *sql.DB

// InitDB initializes the SQLite database
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "vulscanner.db")
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	if DB == nil {
		log.Fatal("Database connection is nil!")
	}

	createTable := `CREATE TABLE IF NOT EXISTS vulnerabilities (
		id TEXT PRIMARY KEY,
		severity TEXT,
		cvss REAL,
		status TEXT,
		package_name TEXT,
		current_version TEXT,
		fixed_version TEXT,
		description TEXT,
		published_date TEXT,
		link TEXT,
		risk_factors TEXT,  
		source_file TEXT,
		scan_time TEXT
	);`

	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}
}

// SaveVulnerability saves a vulnerability record into the database
func SaveVulnerability(vuln models.Vulnerability, sourceFile string) error {
	//  Convert `[]string` to a comma-separated string before saving
	riskFactorsStr := strings.Join(vuln.RiskFactors, ", ")

	_, err := DB.Exec(`INSERT OR IGNORE INTO vulnerabilities 
		(id, severity, cvss, status, package_name, current_version, fixed_version, description, 
		published_date, link, risk_factors, source_file, scan_time) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		vuln.ID, vuln.Severity, vuln.CVSS, vuln.Status, vuln.PackageName,
		vuln.CurrentVersion, vuln.FixedVersion, vuln.Description,
		vuln.PublishedDate, vuln.Link, riskFactorsStr, 		
                sourceFile, time.Now().Format(time.RFC3339))

	return err
}
