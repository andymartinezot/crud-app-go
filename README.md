# CRUD Application in Go

This is a simple CRUD application written in Go. It uses MySQL as the database and Docker Compose for container orchestration.

## Project Structure
    ```
    /Users/andymartinez/go/src/github.com/andymartinezot/crud-app-go/backend/
    ├── cmd
    │ └── server
    │ └── main.go
    ├── config
    │ └── config.go
    ├── internal
    │ ├── database
    │ │ └── database.go
    │ ├── handlers
    │ │ └── handlers.go
    │ ├── models
    │ │ └── employee.go
    │ └── templates
    │ ├── create.tmpl
    │ ├── footer.tmpl
    │ ├── header.tmpl
    │ ├── init.tmpl
    │ └── update.tmpl
    ├── go.mod
    ├── go.sum
    └── Dockerfile
    └── docker-compose.yml
    └── .env
    ```