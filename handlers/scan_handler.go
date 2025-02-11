package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"

	"vulscanner/database"
	"vulscanner/models"
)

// ScanRequest struct to parse incoming JSON request
type ScanRequest struct {
	Repo string `json:"repo"`
}

// ScanResponse struct for API response
type ScanResponse struct {
	Message       string   `json:"message"`
	ProcessedFiles []string `json:"processed_files"`
	Status        string   `json:"status"`
}

// GitHub API URL to fetch repository contents
const githubAPI = "https://api.github.com/repos/%s/contents/"

// ScanRepoHandler handles scanning a GitHub repository for JSON files
func ScanRepoHandler(w http.ResponseWriter, r *http.Request) {
	var req ScanRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Fetch all JSON files dynamically
	jsonFiles, err := fetchJSONFiles(req.Repo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch JSON files: %v", err), http.StatusInternalServerError)
		return
	}

	// Process each JSON file concurrently
	var wg sync.WaitGroup
	for _, file := range jsonFiles {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			processFile(req.Repo, f)
		}(file)
	}
	wg.Wait()

	// Return structured JSON response
	response := ScanResponse{
		Message:       "Scanning initiated",
		ProcessedFiles: jsonFiles,
		Status:        "success",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// fetchJSONFiles retrieves all JSON files from a given GitHub repository
func fetchJSONFiles(repo string) ([]string, error) {
	url := fmt.Sprintf(githubAPI, repo)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch repo contents: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	// Parse the JSON response from GitHub
	var files []map[string]interface{}
	if err := json.Unmarshal(body, &files); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	// Extract JSON file names
	var jsonFiles []string
	for _, file := range files {
		if name, ok := file["name"].(string); ok && strings.HasSuffix(name, ".json") {
			jsonFiles = append(jsonFiles, name)
		}
	}

	return jsonFiles, nil
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
