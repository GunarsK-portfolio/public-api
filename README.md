# Public API

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
FILES_API_URL=http://localhost:8085
```

3. Start infrastructure (if not running):
```bash
# From infrastructure directory
docker-compose up -d postgres files-api flyway
```

4. Run the service:
```bash
task run
# or
go run cmd/api/main.go
```

## Available Commands

Using Task:
```bash
task run           # Run the service
task build         # Build binary
task test          # Run tests
task swagger       # Generate Swagger docs
task clean         # Clean build artifacts
task docker-build  # Build Docker image
task docker-run    # Run with docker-compose
task docker-logs   # View logs
```

Using Go directly:
```bash
go run cmd/api/main.go       # Run
go build -o bin/public-api cmd/api/main.go  # Build
go test ./...                # Test
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
| `FILES_API_URL` | Files API endpoint | `http://localhost:8085` |

## Development

### Running Tests

```bash
task test
# or
go test ./...
```

### Generating Swagger Docs

```bash
task swagger
# or
swag init -g cmd/api/main.go -o docs
```

### Building

```bash
task build
# or
go build -o bin/public-api cmd/api/main.go
```

## Integration

This API is consumed by the public-web frontend to display portfolio content.

## License

MIT
