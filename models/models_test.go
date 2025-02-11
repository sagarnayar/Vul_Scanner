package models

import (
	"encoding/json"
	"testing"
)

// TestVulnerabilityJSONParsing verifies JSON parsing for the Vulnerability model
func TestVulnerabilityJSONParsing(t *testing.T) {
	jsonData := `{
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
		"risk_factors": ["Remote Code Execution", "High CVSS Score"],  
		"source_file": "test.json",
		"scan_time": "2025-02-11T01:44:40-05:00"
	}`

	var vuln Vulnerability
	err := json.Unmarshal([]byte(jsonData), &vuln)
	if err != nil {
		t.Fatalf("Failed to parse JSON: %v", err) 	}

	// Validate `RiskFactors` is correctly parsed as a slice
	if len(vuln.RiskFactors) == 0 {
		t.Errorf("Expected RiskFactors to be a non-empty array, got empty")
	}
}
