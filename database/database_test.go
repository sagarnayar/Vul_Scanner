package database

import (
	"testing"
	"vulscanner/models"
)

func TestSaveVulnerability(t *testing.T) {
	InitDB() // Ensure DB is initialized

	vuln := models.Vulnerability{
		ID:             "CVE-2024-5555",
		Severity:       "HIGH",
		CVSS:           8.5,
		Status:         "active",
		PackageName:    "tensorflow",
		CurrentVersion: "2.7.0",
		FixedVersion:   "2.7.1",
		Description:    "Remote code execution in TensorFlow",
		PublishedDate:  "2025-01-24T00:00:00Z",
		Link:           "https://nvd.nist.gov/vuln/detail/CVE-2024-5555",
		RiskFactors:    []string{"Remote Code Execution", "High CVSS Score"}, 
		SourceFile:     "test.json",
		ScanTime:       "2025-02-11T01:44:40-05:00",
	}

	err := SaveVulnerability(vuln, "test.json")
	if err != nil {
		t.Errorf("Failed to save vulnerability: %v", err)
	}

	rows, err := DB.Query("SELECT id FROM vulnerabilities WHERE id = ?", vuln.ID)
	if err != nil || !rows.Next() {
		t.Errorf("Vulnerability not found in database")
	}
	rows.Close()
}
