# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy source code
COPY . .

# Download dependencies and build
RUN go mod tidy && go mod download
RUN go build -o public-api ./cmd/api

# Production stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/public-api .

EXPOSE 8082

CMD ["./public-api"]
