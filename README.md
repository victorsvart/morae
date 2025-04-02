# Morae

## Overview

Morae is a Golang REST API with minimum usage of external libraries.

## Prerequisites

![Go](https://img.shields.io/badge/Go-1.24+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Latest-336791?style=for-the-badge&logo=postgresql&logoColor=white)
![Make](https://img.shields.io/badge/Make-Required-FF69B4?style=for-the-badge&logo=gnu&logoColor=white)
![direnv](https://img.shields.io/badge/direnv-Required-75D037?style=for-the-badge&logo=vim&logoColor=white)
![Air](https://img.shields.io/badge/Air-Hot_Reload-00BFFF?style=for-the-badge&logo=go&logoColor=white)

## Installation

1. Clone the repository:

   ```
   git clone https://github.com/yourusername/morae.git
   cd morae
   ```

2. Install dependencies:

   ```
   go mod download
   ```

3. Install Air for hot reloading (if not already installed):
   ```
   go install github.com/cosmtrek/air@latest
   ```

## Environment Setup

Morae uses direnv for environment variable management.

### Using `.envrc` file

Create a `.envrc` file in the project root with content similar to:

```
export PORT=:8080
export HOST=localhost
export DSN=postgres://youruser:yourpasword@localhost/databasename
export DB_MAX_OPEN_CONNS=30
export DB_MAX_IDLE_CONNS=30
export DB_MAX_IDLE_TIME=15min
```

Then allow direnv to load the file:

```
direnv allow
```

## Database Setup

1. Create a PostgreSQL database:

   ```
   createdb databasename
   ```

2. Run migrations:
   ```
   make migrate-up
   ```

## Running the Application

### Development Mode with Hot Reloading

```
air
```

### Standard Development Mode

```
make run
```

### Build and Run

```
make build
./morae
```

### Testing

```
make test
```

## Makefile Commands

- `make migration [name]`: Create a new migration file with the specified name
- `make migrate-up`: Run all pending migrations
- `make migrate-down [n]`: Rollback the last n migrations (defaults deletes all migrations, so careful with that)
- `make build`: Compile the application
- `make run`: Run the application in development mode
- `make test`: Run tests
- `make clean`: Clean build artifacts

## Project Structure

```
morae/
├── cmd/
│   └── migrate/
│       └── migrations/  # Database migrations
├── internal/       # Private application code
├── pkg/            # Public libraries
├── configs/        # Configuration files
├── .air.toml       # Air configuration file
├── Makefile        # Build commands
├── go.mod          # Go module definition
├── go.sum          # Go module checksums
└── .envrc          # Environment configuration
```
