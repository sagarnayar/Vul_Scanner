package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"vulscanner/database"
	"vulscanner/models"
)

type QueryRequest struct {
	Filters map[string]string `json:"filters"`
}

// QueryVulnerabilitiesHandler retrieves vulnerabilities based on severity filter
func QueryVulnerabilitiesHandler(w http.ResponseWriter, r *http.Request) {
	var req QueryRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("Invalid JSON:", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	log.Printf("Received query: %+v", req)

	severity, exists := req.Filters["severity"]
	if !exists {
		log.Println("Missing severity filter")
		http.Error(w, "Missing severity filter", http.StatusBadRequest)
		return
	}

	log.Printf("Executing SQL Query: SELECT * FROM vulnerabilities WHERE severity = '%s'", severity)

	rows, err := database.DB.Query(`SELECT id, severity, cvss, status, package_name, current_version, 
		fixed_version, description, published_date, link, risk_factors, source_file, scan_time 
		FROM vulnerabilities WHERE severity = ?`, severity)

	if err != nil {
		log.Println("Database query error:", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var results []models.Vulnerability
	for rows.Next() {
		var vuln models.Vulnerability
		var riskFactorsStr string 

		err = rows.Scan(&vuln.ID, &vuln.Severity, &vuln.CVSS, &vuln.Status, &vuln.PackageName,
			&vuln.CurrentVersion, &vuln.FixedVersion, &vuln.Description, &vuln.PublishedDate,
			&vuln.Link, &riskFactorsStr, &vuln.SourceFile, &vuln.ScanTime)

		if err != nil {
			log.Println("Error scanning row:", err)
			continue
		}

		// Convert riskFactorsStr back to a slice
		vuln.RiskFactors = strings.Split(riskFactorsStr, ", ")

		results = append(results, vuln)
	}

	log.Printf("Final Query Results: %+v", results)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
