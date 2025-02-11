package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestScanRepoHandler(t *testing.T) {
	requestBody := `{"repo": "velancio/vulnerability_scans", "files": ["vulnscan1011.json"]}`
	req, err := http.NewRequest("POST", "/scan", bytes.NewBuffer([]byte(requestBody)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ScanRepoHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status 200, got %d", status)
	}
}
