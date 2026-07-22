VERSION ?= 0.1.0
COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "dev")
DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

LDFLAGS = \
-X github.com/celeguim/emp-cli/internal/version.Version=$(VERSION) \
-X github.com/celeguim/emp-cli/internal/version.GitCommit=$(COMMIT) \
-X github.com/celeguim/emp-cli/internal/version.BuildDate=$(DATE)

build:
	go build -ldflags "$(LDFLAGS)" -o bin/emp
	