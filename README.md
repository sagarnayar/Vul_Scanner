# VulScanner

VulScanner is a **Go-based service** that scans a GitHub repository for JSON files containing vulnerability data, stores them in an **SQLite database**, and allows querying based on severity.

## Features

✅ **Scan API** (POST `/scan`): Fetches and processes JSON vulnerability reports from a GitHub repository.  
✅ **Query API** (POST `/query`): Retrieves stored vulnerabilities filtered by severity.  
✅ **Concurrency**: Processes multiple files in parallel for efficiency.  
✅ **Error Handling**: Retries GitHub API calls (up to **2 times**) upon failure.  
✅ **Docker Support**: Runs as a single-container service.  
✅ **Testing**: Achieves **70%+ test coverage** with unit and integration tests.  

---

## 📌 **Installation & Setup**

### **Prerequisites**
- **Go** (>= 1.19)  
- **SQLite3**  
- **Docker** (optional, for containerized deployment)  

### **Clone the Repository**
```sh
git clone https://github.com/yourusername/vulscanner.git
cd vulscanner
```

### **Install Dependencies**
```sh
go mod tidy
```

### **Build the Project**
```sh
go build -o vulscanner
```

### **Run the Service**
```sh
./vulscanner
```

---

## 📌 **Running & Testing Instructions**

### **Running the Service**
1. Ensure you have installed dependencies.
2. Start the service:
   ```sh
   ./vulscanner
   ```
3. The service will be available at `http://localhost:8080`.

### **Testing Instructions**

#### **Automated Testing**
Run the following command to execute tests and check test coverage:
```sh
go test ./... -cover
```
- Unit tests cover core functionalities, including JSON parsing and database operations.  
- Integration tests validate API behavior with simulated GitHub responses.  
- Test coverage report ensures code reliability and stability.  

#### **Manual Testing**
1. **Run the service** using `./vulscanner`.
2. **Test Scan API:**
   - Use Postman or `curl` to send a POST request to `/scan` with a valid GitHub repo URL.
   - Verify that JSON files are processed and stored in SQLite.
   ```sh
   curl -X POST "http://localhost:8080/scan" -H "Content-Type: application/json" -d '{"repo": "velancio/vulnerability_scans", "files": ["vulnscan1011.json"]}'
   ```
3. **Test Query API:**
   - Send a GET request to `/query` with appropriate filters.
   - Verify that relevant JSON data is returned.
   ```sh
   curl -X POST "http://localhost:8080/query" -H "Content-Type: application/json" -d '{"filters": {"severity": "HIGH"}}'
   ```

### **Example Query Response**
```sh
PS C:\Users\sagar> Invoke-RestMethod -Uri "http://localhost:8080/query" -Method Post -Headers @{"Content-Type"="application/json"} -Body '{"filters": {"severity": "HIGH"}}'
```
**Sample Output:**
```json
{
  "id": "CVE-2024-2222",
  "severity": "HIGH",
  "cvss": 8.2,
  "status": "active",
  "package_name": "spring-security",
  "current_version": "5.6.0",
  "fixed_version": "5.6.1",
  "description": "Authentication bypass in Spring Security",
  "published_date": "2025-01-27T00:00:00Z",
  "link": "https://nvd.nist.gov/vuln/detail/CVE-2024-2222",
  "risk_factors": "Authentication Bypass, High CVSS Score, Proof of Concept Exploit Available",
  "source_file": "vulnscan1011.json",
  "scan_time": "2025-02-11T01:44:40-05:00"
}
```

---

## 📌 **Docker Deployment**

### **Build the Docker Image**
```sh
docker build -t vulscanner .
```

### **Run the Container**
```sh
docker run -p 8080:8080 vulscanner
```

---

