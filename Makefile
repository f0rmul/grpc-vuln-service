build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/main.go

run:
	go run ./cmd/main.go

test:
	go test -v -count=1 ./...
start: build
	docker-compose up --build vuln-service

deps-reset:
	git checkout -- go.mod
	go mod tidy
	go mod vendor

tidy:
	go mod tidy
	go mod vendor

lint:
	echo "Starting linters"
	golangci-lint run ./...

PHONY: generate
generate:
	mkdir -p pkg/netvuln_v1
	protoc --go_out=./pkg/netvuln_v1 --go_opt=paths=source_relative \
	       --go-grpc_out=./pkg/netvuln_v1 --go-grpc_opt=paths=source_relative \
		   api/netvuln_v1/netvuln.proto  
	mv pkg/netvuln_v1/api/netvuln_v1/* pkg/netvuln_v1
	rm -rf pkg/netvuln_v1/api