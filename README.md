# Portfolio Public API

Read-only REST API for serving public portfolio data. Provides endpoints for profile information, work experience, certifications, and miniature painting projects.

## Features

- Read-only REST API
- CORS enabled for public access
- Swagger API documentation
- S3/MinIO integration for image URLs
- Docker support
- GORM with PostgreSQL

## API Endpoints

### Profile
- `GET /api/v1/profile` - Get profile information

### Experience
- `GET /api/v1/experience` - List all work experience

### Certifications
- `GET /api/v1/certifications` - List all certifications

### Miniatures
- `GET /api/v1/miniatures` - List all miniature projects (with images)
- `GET /api/v1/miniatures/:id` - Get specific miniature project details

### Health
- `GET /api/v1/health` - Service health check

### Documentation
- `GET /swagger/index.html` - Swagger UI

## Quick Start

### Prerequisites

- Go 1.25+
- PostgreSQL 18+
- MinIO or S3
- [Task](https://taskfile.dev/installation/) (task runner)
- Docker (optional)

### Local Development (without Docker)

```bash
# Install dependencies
go mod download

# Copy environment file
cp .env.example .env

# Edit .env with your local settings
# (Defaults should work if you have postgres/minio running locally)

# Generate Swagger docs
task swagger

# Run the service
task run

# Or debug in VS Code (F5)
```

### Local Development (with Docker)

```bash
# Start all services (PostgreSQL, MinIO, Public API)
docker-compose up -d

# View logs
docker-compose logs -f public-api

# Stop services
docker-compose down
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| PORT | Service port | 8082 |
| DB_HOST | PostgreSQL host | localhost |
| DB_PORT | PostgreSQL port | 5432 |
| DB_USER | Database user | portfolio_user |
| DB_PASSWORD | Database password | portfolio_pass |
| DB_NAME | Database name | portfolio |
| S3_ENDPOINT | S3/MinIO endpoint | http://localhost:9000 |
| S3_ACCESS_KEY | S3 access key | minioadmin |
| S3_SECRET_KEY | S3 secret key | minioadmin |
| S3_BUCKET | S3 bucket name | images |
| S3_USE_SSL | Use SSL for S3 | false |

## Project Structure

```
public-api/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go            # Configuration
│   ├── database/
│   │   └── database.go          # Database connection
│   ├── handlers/
│   │   └── handlers.go          # HTTP handlers
│   ├── models/
│   │   └── models.go            # Data models
│   ├── repository/
│   │   └── repository.go        # Database queries
│   └── storage/
│       └── s3.go                # S3 client (future)
├── docs/                        # Swagger documentation
├── Dockerfile
├── docker-compose.yml
├── Taskfile.yml
├── go.mod
└── README.md
```

## API Usage Examples

### Get Profile

```bash
# Direct access (standalone)
curl http://localhost:8082/api/v1/profile

# Via Traefik (infrastructure setup)
curl http://localhost/api/v1/profile
```

Response:
```json
{
  "id": 1,
  "full_name": "Your Name",
  "title": "Software Engineer",
  "bio": "Passionate about software development...",
  "email": "your.email@example.com",
  "phone": "+1234567890",
  "location": "Your City, Country",
  "avatar_url": "https://...",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

### Get Work Experience

```bash
# Direct access (standalone)
curl http://localhost:8082/api/v1/experience

# Via Traefik (infrastructure setup)
curl http://localhost/api/v1/experience
```

Response:
```json
[
  {
    "id": 1,
    "company": "Example Company",
    "position": "Senior Developer",
    "description": "Led development of key features...",
    "start_date": "2020-01-01",
    "end_date": null,
    "is_current": true,
    "display_order": 1,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
]
```

### Get Miniature Projects

```bash
# Direct access (standalone)
curl http://localhost:8082/api/v1/miniatures

# Via Traefik (infrastructure setup)
curl http://localhost/api/v1/miniatures
```

Response:
```json
[
  {
    "id": 1,
    "title": "Space Marine Squad",
    "description": "Painted a full squad...",
    "completed_date": "2024-01-15",
    "display_order": 1,
    "images": [
      {
        "id": 1,
        "miniature_project_id": 1,
        "title": "Front view",
        "description": "Squad front view",
        "url": "https://...",
        "display_order": 1,
        "created_at": "2024-01-01T00:00:00Z",
        "updated_at": "2024-01-01T00:00:00Z"
      }
    ],
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
]
```

### Get Specific Miniature Project

```bash
# Direct access (standalone)
curl http://localhost:8082/api/v1/miniatures/1

# Via Traefik (infrastructure setup)
curl http://localhost/api/v1/miniatures/1
```

## Development

### Generate Swagger Documentation

```bash
# Install swag (one-time)
go install github.com/swaggo/swag/cmd/swag@latest

# Generate docs
task swagger

# Commit generated docs
git add docs/
git commit -m "Update swagger docs"
```

### Run Tests

```bash
task test
```

### Build Binary

```bash
task build
```

## Docker

### Build Image

```bash
task docker-build
```

### Run with Docker Compose

```bash
task docker-run
```

### View Logs

```bash
task docker-logs
```

## CORS Configuration

The API has CORS enabled to allow cross-origin requests from the public web frontend:

```go
Access-Control-Allow-Origin: *
Access-Control-Allow-Methods: GET, OPTIONS
Access-Control-Allow-Headers: Content-Type
```

For production, restrict the origin to your actual domain.

## Database

This service uses the shared portfolio database. Make sure migrations have been run:

```bash
cd ../database
docker-compose run flyway migrate
```

## Troubleshooting

### Database Connection Failed

```bash
# Check PostgreSQL is running
docker-compose ps postgres

# Test connection
psql postgresql://portfolio_user:portfolio_pass@localhost:5432/portfolio
```

### MinIO Connection Issues

```bash
# Check MinIO is running
docker-compose ps minio

# Access MinIO console
open http://localhost:9001
```

### Port Already in Use

```bash
# Find process using port 8082
lsof -i :8082

# Kill process
kill -9 <PID>
```

## Related Repositories

- [infrastructure](https://github.com/GunarsK-portfolio/infrastructure)
- [database](https://github.com/GunarsK-portfolio/database)
- [public-web](https://github.com/GunarsK-portfolio/public-web)
- [auth-service](https://github.com/GunarsK-portfolio/auth-service)
- [admin-api](https://github.com/GunarsK-portfolio/admin-api)

## License

MIT
