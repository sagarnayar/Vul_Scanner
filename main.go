package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"vulscanner/database"
	"vulscanner/handlers"
)

func main() {
	database.InitDB()

	r := mux.NewRouter()
	r.HandleFunc("/scan", handlers.ScanRepoHandler).Methods("POST")
	r.HandleFunc("/query", handlers.QueryVulnerabilitiesHandler).Methods("POST")

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
