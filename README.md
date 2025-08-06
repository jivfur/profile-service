# Profile Service

Profile Service is a backend microservice responsible for managing user profiles within the Out&Proud ecosystem. It handles user registration, profile updates, face and email verification statuses, and provides verified profile data to other services.

---

## Features

- User profile CRUD (Create, Read, Update, Delete) operations  
- Face photo verification integration (planned)  
- Email verification enforcement (planned)  
- Secure and privacy-conscious design  
- RESTful API built with Go and GORM  
- Data stored in MySQL with proper relational structure  
- Structured for maintainability and scalability  

---

## Technology Stack

- Language: Go (Golang)  
- ORM: GORM  
- Database: MySQL  
- Containerization: Docker (planned)  
- Orchestration: Kubernetes (planned)  

---

## Getting Started

### Prerequisites

- Go 1.20+  
- MySQL 8+  
- Docker (optional, for containerized development)  

### Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/YourOrg/profile-service.git
   cd profile-service
   ```
2. Configure your database connection in config/ (environment variables recommended).

3. Run database migrations (using GORM AutoMigrate or migration tool, TBD).

4. Start the service:
   ```bash
   go run cmd/server/main.go
   ```
### API Endpoints (Examples)
- POST /profiles — Create a new user profile
- GET /profiles/{id} — Retrieve profile information
- PUT /profiles/{id} — Update profile details
- POST /profiles/{id}/photos — Upload profile photos
- GET /profiles/search — Search for verified profiles

### Project Structure
- cmd/ — Application entrypoint(s)
- internal/handler/ — HTTP handlers for REST API
- internal/service/ — Business logic layer
- internal/repository/ — Database access layer
- internal/model/ — Data models and schema
- config/ — Configuration and environment setup


