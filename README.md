# VulScanner

VulScanner is a **Go-based service** that scans a GitHub repository for JSON files containing vulnerability data, stores them in an **SQLite database**, and allows querying based on severity.

## Features

✅ **Scan API** (POST `/scan`): Fetches and processes JSON vulnerability reports from a GitHub repository.  
✅ **Query API** (POST `/query`): Retrieves stored vulnerabilities filtered by severity.  
✅ **Concurrency**: Processes multiple files in parallel for efficiency.  
✅ **Error Handling**: Retries GitHub API calls (up to **2 times**) upon failure.  
✅ **Docker Support**: Runs as a single-container service.  
✅ **Testing**: Achieves **70%+ test coverage**.  

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

Install Dependencies
sh
Copy
Edit
go mod tidy
Build the Projec
