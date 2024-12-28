# Base stage
FROM golang:1.23.4-alpine AS base

RUN apk add --no-cache git
WORKDIR /app

# Copy dependencies
COPY go.* ./
RUN go mod download

# Copy source code
COPY . .

# Development stage
FROM base AS development

# Install air for hot reloading
RUN go install github.com/air-verse/air@latest

WORKDIR /app/cmd/report

# Use air for development
CMD ["air", "-c", ".air.toml"]

# Production stage
FROM base AS production

RUN go build -o report_app cmd/report/main.go

# Runner stage
FROM alpine:latest AS runner

WORKDIR /app

COPY --from=production /app/report_app .

RUN chmod +x ./report_app

# Command to run the production binary
CMD ["./report_app"]
