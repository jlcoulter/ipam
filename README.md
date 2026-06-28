# Go API Service Template

A GitHub template for Go API services — HTTP server, structured logging, health checks, graceful shutdown, Docker, CI, and release automation included.

## Features

- **HTTP server** with [chi](https://github.com/go-chi/chi) router
- **Structured logging** with [slog](https://pkg.go.dev/log/slog) (stdlib)
- **Health & readiness** endpoints at `/healthz` and `/readyz`
- **Configuration** via environment variables with [viper](https://github.com/spf13/viper)
- **Graceful shutdown** on SIGINT/SIGTERM
- **Docker** multi-stage build (scratch image, ~10MB)
- **CI** via GitHub Actions (test, vet, build on push/PR)
- **Release** via [GoReleaser](https://goreleaser.com/) on tag push
- **Makefile** for common tasks

## Usage

1. Click **"Use this template"** on GitHub to create a new repo
2. Run the setup script:
   ```sh
   ./setup.sh my-service github.com/you/my-service
   ```
3. Start building your handlers in `internal/handler/`
4. Add routes in `cmd/server/main.go`

## Project Structure

```
.
├── cmd/
│   └── server/
│       └── main.go          # Entry point, router setup, graceful shutdown
├── internal/
│   ├── config/
│   │   └── config.go        # Viper config (env vars + defaults)
│   ├── handler/
│   │   ├── health.go        # /healthz, /readyz
│   │   └── example.go       # Example CRUD handlers (delete me)
│   ├── middleware/
│   │   └── logging.go       # Request logging middleware
│   └── store/
│       └── store.go         # Database interface (swap implementations)
├── .github/
│   └── workflows/
│       └── ci.yml            # Test + vet + build
├── .goreleaser.yml
├── Dockerfile
├── Makefile
├── go.mod
├── setup.sh                  # Template setup script (rename module, remove examples)
└── README.md
```

## Quick Start

```sh
# Run locally
make run

# Run tests
make test

# Build binary
make build

# Build Docker image
make docker
```

## Container Images

CI builds and pushes a container image to GHCR on every push to any branch.

```sh
# Pull the latest image
docker pull ghcr.io/<owner>/ipam:latest

# Pull a specific commit
docker pull ghcr.io/<owner>/ipam:<sha>

# Run
docker run -p 8080:8080 ghcr.io/<owner>/ipam:latest
```

Replace `<owner>` with your GitHub username or org. Images are tagged with both `latest` and the commit SHA.

## Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | HTTP listen port |
| `LOG_LEVEL` | `info` | Log level (debug, info, warn, error) |
| `DB_URI` | "" | Database connection string (optional) |

## Make Targets

| Target | Description |
|--------|-------------|
| `make run` | Run the server locally |
| `make build` | Build the binary to `./bin/` |
| `make test` | Run tests |
| `make vet` | Run `go vet` |
| `make lint` | Run golangci-lint |
| `make docker` | Build Docker image |
| `make clean` | Remove build artifacts |

## License

MIT