run:
	@go run cli/main.go run -p ./pkg/world/providers/testdata/world_map.txt -n 10 -m 10

build:
	@go build ./...

test:
	@go test -v ./...