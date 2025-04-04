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
   go mod tidy
   ```

3. Install Air for hot reloading (if not already installed). This is not necessary, check Make scripts bellow to run without hot reloading:
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

Install direnv on your Linux distribution:

```bash
# Ubuntu/Debian
sudo apt update
sudo apt install direnv

# Fedora
sudo dnf install direnv

# Arch Linux
sudo pacman -S direnv

# CentOS/RHEL (using EPEL repository)
sudo yum install epel-release
sudo yum install direnv

# openSUSE
sudo zypper install direnv

# Gentoo
sudo emerge app-shells/direnv

# Alpine Linux
sudo apk add direnv

# Using Homebrew (cross-platform)
brew install direnv
```

## Shell Configuration

Add the appropriate hook to your shell configuration file:

```bash
# For bash (~/.bashrc)
eval "$(direnv hook bash)"

# For zsh (~/.zshrc)
eval "$(direnv hook zsh)"

# For fish (~/.config/fish/config.fish)
direnv hook fish | source

# For tcsh (~/.tcshrc)
eval `direnv hook tcsh`
```

## Usage

Allow your .envrc file using direnv in the root directory:

```bash
direnv allow
```

## Database Setup

1. Create a PostgreSQL database:

   ```
   $ createdb databasename
   ```

2. Run migrations:
   ```
   $ make migrate-up
   ```

## Running the Application

### Development Mode with Hot Reloading

```
$ air
```

### Standard Development Mode

```
$ make run
```

### Build and Run

```
$ make build
$ ./morae
```

### Testing

```
$ make test
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
