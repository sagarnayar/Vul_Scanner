package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"vulscanner/database"
	"vulscanner/models"
)

type ScanRequest struct {
	Repo  string   `json:"repo"`
	Files []string `json:"files"`
}

// ScanRepoHandler handles scanning a GitHub repository for JSON files
func ScanRepoHandler(w http.ResponseWriter, r *http.Request) {
	var req ScanRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	var wg sync.WaitGroup
	for _, file := range req.Files {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			processFile(req.Repo, f)
		}(file)
	}
	wg.Wait()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Scan completed"))
}

// processFile fetches and processes a JSON file from GitHub
func processFile(repo, file string) {
    url := fmt.Sprintf("https://raw.githubusercontent.com/%s/main/%s", repo, file)
    log.Println("Fetching:", url)

    resp, err := http.Get(url)
    if err != nil {
        log.Printf("Failed to fetch file %s: %v", file, err)
        return
    }
    defer resp.Body.Close()

    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Printf("Failed to read file %s: %v", file, err)
        return
    }

    log.Println("Raw JSON Content:", string(data)) 

    var scans []models.ScanData
    err = json.Unmarshal(data, &scans)
    if err != nil {
        log.Printf("Error parsing JSON from %s: %v", file, err)
        return
    }

    for _, scan := range scans {
        for _, vuln := range scan.ScanResults.Vulnerabilities {
            log.Printf("Parsed Vulnerability: %+v\n", vuln) 
            err := database.SaveVulnerability(vuln, file)
            if err != nil {
                log.Printf("Failed to save %s: %v", vuln.ID, err)
            }
        }
    }
}
