# Testing Guide

## Overview

The public-api uses Go's standard `testing` package with httptest for handler
unit tests. This is a read-only API serving public portfolio data.

## Quick Commands

```bash
# Run all tests
go test ./internal/handlers/

# Run with coverage
go test -cover ./internal/handlers/

# Generate coverage report
go test -coverprofile=coverage.out ./internal/handlers/
go tool cover -html=coverage.out -o coverage.html

# Run specific test
go test -v -run TestGetProfile_Success ./internal/handlers/

# Run all Project tests
go test -v -run Project ./internal/handlers/

# Run all Miniature tests
go test -v -run Miniature ./internal/handlers/
```

## Test Files

**`handler_test.go`** - 36 tests

| Category | Tests | Coverage |
| -------- | ----- | -------- |
| Profile | 4 | GetProfile + error cases |
| Work Experience | 3 | GetAll + error cases |
| Certifications | 3 | GetAll + error cases |
| Skills | 2 | GetAll + error cases |
| Projects | 7 | GetAll, GetByID + error cases |
| Miniatures | 7 | GetAll, GetByID + error cases |
| Miniature Themes | 7 | GetAll, GetByID + error cases |
| Context Propagation | 1 | Verifies context with sentinel value |
| ID Validation | 1 | Table-driven invalid ID format tests |
| Constructor | 1 | Handler initialization |

## Key Testing Patterns

**Mock Repository**: Function fields allow per-test behavior customization

```go
mockRepo.getProfileFunc = func(ctx context.Context) (*models.Profile, error) {
    return &expectedProfile, nil
}
```

**HTTP Testing**: Uses `httptest.ResponseRecorder` with Gin router

```go
w := performRequest(router, "GET", "/profile", nil)
if w.Code != http.StatusOK { ... }
```

**Test Helpers**: Factory functions for consistent test data

```go
profile := createTestProfile()
project := createTestProject()
miniature := createTestMiniatureProject()
```

## Test Categories

### Success Cases

- Returns expected data
- Sets correct HTTP status (200 OK)
- Empty arrays returned for no data

### Error Cases

- Repository errors (500 Internal Server Error)
- Not found errors (404 Not Found) - for GetByID endpoints
- Invalid ID format (400 Bad Request)

## API Characteristics

Public-api is **read-only**:

- No authentication required
- No Create/Update/Delete operations
- Database user has SELECT-only permissions

## Contributing Tests

1. Follow naming: `Test<HandlerName>_<Scenario>`
2. Organize by endpoint with section markers
3. Mock only the repository methods needed
4. Verify: `go test -cover ./internal/handlers/`
