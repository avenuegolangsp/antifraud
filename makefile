.PHONY: build build-web build-engine
build:
	@mkdir -p bin
	@echo "Building web service..."
	@cd internal/services/web && go build -o ../../../bin/antifraud-webserver cmd/main.go
	@echo "Building engine service..."
	@cd internal/services/engine && go build -o ../../../bin/antifraud-engine cmd/main.go

build-web:
	@mkdir -p bin
	@echo "Building web service..."
	@cd internal/services/web && go build -o ../../../bin/antifraud-webserver cmd/main.go

build-engine:
	@mkdir -p bin
	@echo "Building engine service..."
	@cd internal/services/engine && go build -o ../../../bin/antifraud-engine cmd/main.go

.PHONY: run run-web run-engine
run:
	@echo "Starting both services..."
	@(cd internal/services/web && go run cmd/main.go) & (cd internal/services/engine && go run cmd/main.go) & wait

run-web:
	cd internal/services/web && go run cmd/main.go

run-engine:
	cd internal/services/engine && go run cmd/main.go

.PHONY: test test-web test-engine
test:
	@echo "Running tests for web service..."
	@cd internal/services/web && go test ./...
	@echo "Running tests for engine service..."
	@cd internal/services/engine && go test ./...

test-web:
	@echo "Running tests for web service..."
	@cd internal/services/web && go test ./...

test-engine:
	@echo "Running tests for engine service..."
	@cd internal/services/engine && go test ./...

# test-integration:
#     go test ./tests/integration/...

.PHONY: test-coverage
test-coverage:
	@echo "Running coverage for web service..."
	@cd internal/services/web && go test -coverprofile=coverage.out ./...
	@cd internal/services/web && go tool cover -html=coverage.out
	@echo "Running coverage for engine service..."
	@cd internal/services/engine && go test -coverprofile=coverage.out ./...
	@cd internal/services/engine && go tool cover -html=coverage.out

.PHONY: lint
lint:
	@echo "Linting web service..."
	@cd internal/services/web && golangci-lint run
	@echo "Linting engine service..."
	@cd internal/services/engine && golangci-lint run

# # Docker
# docker-build:
#     docker build -t trading-system .

# docker-run:
#     docker-compose up -d
# # Limpeza

.PHONY: clean
clean:
	rm -rf bin/

.PHONY: gomodtidy
gomodtidy:
	@echo "==> Running go mod tidy recursively"
	@find . -name go.mod -execdir sh -c 'pwd && go mod tidy' \;
