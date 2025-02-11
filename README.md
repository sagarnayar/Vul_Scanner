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

## 📌 **Usage**

### **1️⃣ Scan GitHub Repository**
```sh
POST /scan
```
**Request Body:**
```json
{
  "repo_url": "https://github.com/example/repo"
}
```
**Response:**
```json
{
  "message": "Scan started successfully"
}
```

### **2️⃣ Query JSON Data**
```sh
GET /query?key=value
```
**Response:**
```json
{
  "results": [
    { "id": 1, "data": "..." }
  ]
}
```

---

## 📌 **Testing Instructions**

### **Automated Testing**
Run the following command to execute tests and check test coverage:
```sh
go test ./... -cover
```
- Unit tests cover core functionalities, including JSON parsing and database operations.  
- Integration tests validate API behavior with simulated GitHub responses.  
- Test coverage report ensures code reliability and stability.  

### **Manual Testing**
1. **Run the service** using `./vulscanner`.
2. **Test Scan API:**
   - Use Postman or `curl` to send a POST request to `/scan` with a valid GitHub repo URL.
   - Verify that JSON files are processed and stored in SQLite.
3. **Test Query API:**
   - Send a GET request to `/query` with appropriate filters.
   - Verify that relevant JSON data is returned.

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
