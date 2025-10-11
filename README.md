# Public API

RESTful API for public portfolio content access.

## Features

- Read-only public portfolio content
- Projects, skills, experience endpoints
- Image serving via MinIO/S3
- RESTful API with Swagger documentation
- Health check endpoint

## Tech Stack

- **Language**: Go 1.25
- **Framework**: Gin
- **Database**: PostgreSQL (GORM)
- **Storage**: MinIO (S3-compatible)
- **Documentation**: Swagger/OpenAPI

## Prerequisites

- Go 1.25+
- PostgreSQL (or use Docker Compose)
- MinIO (or use Docker Compose)

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
│   ├── repository/       # Data access layer
│   ├── service/          # Business logic
│   └── storage/          # S3/MinIO integration
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
DB_USER=portfolio_user
DB_PASSWORD=portfolio_pass
DB_NAME=portfolio
S3_ENDPOINT=http://localhost:9000
S3_ACCESS_KEY=minioadmin
S3_SECRET_KEY=minioadmin
S3_BUCKET=images
S3_USE_SSL=false
```

3. Start infrastructure (if not running):
```bash
# From infrastructure directory
docker-compose up -d postgres minio flyway
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
- `GET /projects` - List all projects
- `GET /projects/:id` - Get project details
- `GET /skills` - List all skills
- `GET /experience` - List work experience
- `GET /about` - Get about information

## Swagger Documentation

When running, Swagger UI is available at:
- `http://localhost:8082/swagger/index.html`

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | `8082` |
| `DB_HOST` | PostgreSQL host | `localhost` |
| `DB_PORT` | PostgreSQL port | `5432` |
| `DB_USER` | Database user | `portfolio_user` |
| `DB_PASSWORD` | Database password | `portfolio_pass` |
| `DB_NAME` | Database name | `portfolio` |
| `S3_ENDPOINT` | MinIO/S3 endpoint | `http://localhost:9000` |
| `S3_ACCESS_KEY` | MinIO access key | `minioadmin` |
| `S3_SECRET_KEY` | MinIO secret key | `minioadmin` |
| `S3_BUCKET` | S3 bucket name | `images` |
| `S3_USE_SSL` | Use SSL for S3 | `false` |

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
