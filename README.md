# Secure-Config Service

A lightweight Go web service that demonstrates secure configuration delivery using the Azure SDK for Go (`aztables`).
Built as part of a HashiCorp project series, this phase lays the foundation for Vault and Boundary integration in the next project.

## Overview
- Written in **Go** with the **Gin** web framework
- Connects to **Azure Table Storage** (via `aztables`) for configuration data
- Uses environment variables for credentials (Vault will replace this later)
- Provides simple health and Azure test routes

## Run Locally

### Prerequisites
- Go 1.25+
- Azure Storage Account
- Azure Table named `ConfigTable` (or any test name)

### Setup
1. Clone the repo:
   ```
   git clone <your-repo-url>
   cd secure-config
   ```

2. Create a .env file with your credentials:
AZURE_STORAGE_ACCOUNT=<your_account>
AZURE_STORAGE_KEY=<your_key>

3. Run the service:
```go run main.go```

4. Test endpoints:
```
curl http://localhost:8080/health
curl http://localhost:8080/azure/test
```
**Expected output:**
{"status":"ok"}

### Environment Variables
AZURE_STORAGE_ACCOUNT: Name of the Azure Storage Account
AZURE_STORAGE_KEY: SHared key for authentication

Note: These will later be sourced from Vault in Project 4.

## Next Steps
- Replace .env secrets with Vault dynamic secrets
- Use Boundary to restrict access to the service runtime
- Add minimal integration tests for ```/azure/test```
