package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"vulscanner/database"
	"vulscanner/models"
)

func TestQueryVulnerabilitiesHandler(t *testing.T) {
	database.InitDB() // Ensure DB is initialized

	vuln := models.Vulnerability{
		ID:             "CVE-2024-2222",
		Severity:       "HIGH",
		CVSS:           8.2,
		Status:         "active",
		PackageName:    "spring-security",
		CurrentVersion: "5.6.0",
		FixedVersion:   "5.6.1",
		Description:    "Authentication bypass in Spring Security",
		PublishedDate:  "2025-01-27T00:00:00Z",
		Link:           "https://nvd.nist.gov/vuln/detail/CVE-2024-2222",
		RiskFactors:    []string{"Authentication Bypass", "High CVSS Score"}, 
		SourceFile:     "mock.json",
		ScanTime:       "2025-02-11T01:44:40-05:00",
	}

	//  Insert test data into the database before querying
	err := database.SaveVulnerability(vuln, "mock.json")
	if err != nil {
		t.Fatalf("Failed to insert mock data: %v", err)
	}

	requestBody := `{"filters": {"severity": "HIGH"}}`
	req, err := http.NewRequest("POST", "/query", bytes.NewBuffer([]byte(requestBody)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(QueryVulnerabilitiesHandler)
	handler.ServeHTTP(rr, req)

	var results []models.Vulnerability
	err = json.Unmarshal(rr.Body.Bytes(), &results)
	if err != nil || len(results) == 0 {
		t.Errorf("Query returned no results, expected at least one")
	}
}
