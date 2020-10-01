run.server:
	go run cmd/server/main.go start

build.server:
	go build -o out/server cmd/server/main.go

build.cli:
	go build -o out/migrator cmd/cli/main.go

test:
	go test ./... -v
