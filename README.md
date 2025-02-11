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

## 📌 **License**
MIT License

## 📌 **Author**
- Your Name (@your-username)
