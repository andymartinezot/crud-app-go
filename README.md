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

```
CREATE DATABASE IF NOT EXISTS system;
USE system;
CREATE TABLE IF NOT EXISTS employees (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL
);
```

## Building and Running with Docker Compose
To build and run the application with Docker Compose, follow these steps:

1. Build and start the services:
```
    docker-compose up --build
```

2. Build and start the services:
```
    docker-compose up
```

You should see output indicating that both the db and app services are running.

3. Access the application:
```
    Open your web browser and navigate to http://localhost:8000.
```

## Directory Structure
    - cmd/server: Contains the main application entry point.
    - onfig: Configuration and database connection settings.
    - internal: Contains the application logic, handlers, models, and templates.
    - Dockerfile: Instructions to build the Docker image.
    - docker-compose.yml: Docker Compose configuration file.
    - init.sql: SQL script to initialize the database.
    - .env: Environment variables for database connection.

## Logging
The application logs important events such as new user additions, updates, and deletions. These logs are printed to the console and can be viewed using docker-compose logs app.

## Troubleshooting
- Check Docker Compose Logs:
```
    docker-compose logs app
```

- Rebuild Docker Images:
```
    docker-compose down
    docker-compose build --no-cache
    docker-compose up
```

- Verify Environment Variables:
```
    Ensure the .env file is present and correctly configured.
```