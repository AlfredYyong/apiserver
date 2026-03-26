# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a lightweight RESTful API server built with Go and the Gin framework. The server provides health check and system monitoring endpoints.

### Key Features
- Built with Gin Web framework for high-performance HTTP services
- Configuration management using Viper (YAML config files and environment variables)
- Log rotation functionality using lumberjack
- Health check endpoints for service, disk, CPU, and memory monitoring
- Security HTTP headers
- CORS support
- Configuration hot reloading

## Project Structure

```
apiserver/
├── conf/          # Configuration files
├── config/        # Configuration management code
├── handler/       # Request handlers
├── router/        # Router configuration
├── main.go        # Application entry point
└── go.mod         # Go module definition
```

## Common Development Tasks

### Building and Running
1. Install dependencies: `go mod tidy`
2. Build: `go build -o apiserver`
3. Run: `./apiserver`

### Configuration
- Default config file: `conf/config.yaml`
- Override with command line: `./apiserver -c /path/to/config.yaml`
- Environment variables with prefix `APISERVER_` (e.g., `APISERVER_RUNMODE=release`)

### Adding New API Endpoints
1. Create handler functions in `handler/` directory
2. Add route configuration in `router/router.go`

## Architecture Overview

The application follows a layered architecture:
1. `main.go` - Entry point that initializes config and starts the HTTP server
2. `config/` - Handles configuration loading and log initialization
3. `router/` - Sets up routes and middleware
4. `handler/` - Contains endpoint implementations
5. Middleware in `router/middleware/` handles cross-cutting concerns like security headers and caching

The server uses Viper for configuration management with automatic environment variable binding and fsnotify for hot reloading. Logging uses both the lexkong/log package and slog with lumberjack for file rotation.