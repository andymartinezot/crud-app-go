# CRUD Application in Go

This is a simple CRUD application written in Go. It uses MySQL as the database and Docker Compose for container orchestration.

## Project Structure
```
    /Users/andymartinez/go/src/github.com/andymartinezot/crud-app-go/backend/
    ├── cmd
    │   └── server
    │       └── main.go
    ├── config
    │   └── config.go
    ├── internal
    │   ├── database
    │   │   └── database.go
    │   ├── handlers
    │   │   └── handlers.go
    │   ├── models
    │   │   └── employee.go
    │   └── templates
    │       ├── create.tmpl
    │       ├── footer.tmpl
    │       ├── header.tmpl
    │       ├── init.tmpl
    │       └── update.tmpl
    ├── go.mod
    ├── go.sum
    ├── Dockerfile
    ├── docker-compose.yml
    ├── init.sql
    └── .env
```

## Prerequisites

- Docker
- Docker Compose

## Environment Variables

Create a `.env` file in the project root directory with the following content:

```env
DB_HOST=db
DB_PORT=3306
DB_USER=root
DB_PASSWORD=password123
DB_NAME=system
MYSQL_ROOT_PASSWORD=password123
MYSQL_DATABASE=system
```

## Initialization Script

Create an init.sql file in the project root directory with the following content to ensure the employees table exists:

CREATE DATABASE IF NOT EXISTS system;
USE system;

```
CREATE TABLE IF NOT EXISTS employees (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL
);
```