# Public API

![CI](https://github.com/GunarsK-portfolio/public-api/workflows/CI/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/GunarsK-portfolio/public-api)](https://goreportcard.com/report/github.com/GunarsK-portfolio/public-api)
[![codecov](https://codecov.io/gh/GunarsK-portfolio/public-api/branch/main/graph/badge.svg)](https://codecov.io/gh/GunarsK-portfolio/public-api)

RESTful API for public portfolio content access.

## Features

- Read-only public portfolio content
- Projects, skills, experience, profile endpoints
- File serving via Files API
- RESTful API with Swagger documentation
- Health check endpoint

## Tech Stack

- **Language**: Go 1.25
- **Framework**: Gin
- **Database**: PostgreSQL (GORM)
- **Storage**: Files API (for images/documents)
- **Documentation**: Swagger/OpenAPI

## Prerequisites

- Go 1.25+
- PostgreSQL (or use Docker Compose)
- Files API running (or use Docker Compose)

## Project Structure

```
public-api/
├── cmd/
│   └── api/              # Application entrypoint
├── internal/
│   ├── config/           # Configuration
│   ├── database/         # Database connection
│   ├── handlers/         # HTTP handlers
│   ├── models/           # Data models
│   └── repository/       # Data access layer
└── docs/                 # Swagger documentation
```

## Quick Start

### Using Docker Compose

```bash
docker-compose up -d
```

### Local Development

1. Copy environment file:
```bash
cp .env.example .env
```

2. Update `.env` with your configuration:
```env
PORT=8082
DB_HOST=localhost
DB_PORT=5432
DB_USER=portfolio_public
DB_PASSWORD=portfolio_public_dev_pass
DB_NAME=portfolio
FILES_API_URL=http://localhost:8085/api/v1
```

3. Start infrastructure (if not running):
```bash
# From infrastructure directory
docker-compose up -d postgres flyway
```

4. Run the service:
```bash
go run cmd/api/main.go
```

## Available Commands

Using Task:
```bash
# Development
task dev:swagger         # Generate Swagger documentation
task dev:install-tools   # Install dev tools (golangci-lint, govulncheck, etc.)

# Build and run
task build               # Build binary
task test                # Run tests
task test:coverage       # Run tests with coverage report
task clean               # Clean build artifacts

# Code quality
task format              # Format code with gofmt
task tidy                # Tidy and verify go.mod
task lint                # Run golangci-lint
task vet                 # Run go vet

# Security
task security:scan       # Run gosec security scanner
task security:vuln       # Check for vulnerabilities with govulncheck

# Docker
task docker:build        # Build Docker image
task docker:run          # Run service in Docker container
task docker:stop         # Stop running Docker container
task docker:logs         # View Docker container logs

# CI/CD
task ci:all              # Run all CI checks (format, tidy, lint, vet, test, vuln)
```

Using Go directly:
```bash
go run cmd/api/main.go                       # Run
go build -o bin/public-api cmd/api/main.go   # Build
go test ./...                                 # Test
```

## API Endpoints

Base URL: `http://localhost:8082/api/v1`

### Health Check
- `GET /health` - Service health status

### Public Endpoints
- `GET /profile` - Get profile information
- `GET /projects` - List all projects
- `GET /projects/:id` - Get project details
- `GET /skills` - List all skills grouped by type
- `GET /experience` - List work experience
- `GET /certifications` - List certifications
- `GET /miniatures/themes` - List miniature painting themes
- `GET /miniatures/projects` - List all miniature projects
- `GET /miniatures/projects/:id` - Get miniature project details

## Swagger Documentation

When running, Swagger UI is available at:
- `http://localhost:8082/swagger/index.html`

## Environment Variables

| Variable | Description | Example |
|----------|-------------|---------|
| `PORT` | Server port | `8082` |
| `DB_HOST` | PostgreSQL host | `localhost` |
| `DB_PORT` | PostgreSQL port | `5432` |
| `DB_USER` | Database user (read-only) | `portfolio_public` |
| `DB_PASSWORD` | Database password | `portfolio_public_dev_pass` |
| `DB_NAME` | Database name | `portfolio` |
| `FILES_API_URL` | Files API base URL | `http://localhost:8085/api/v1` |

## Integration

This API is consumed by the public-web frontend to display portfolio content.

## License

MIT
