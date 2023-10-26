build:
	go build -ldflags="-s -w" -o ./bin/exe ./cmd/cli

s:
	go run cmd/cli/main.go s

w:
	go run cmd/cli/main.go w

j:
	go run cmd/cli/main.go j
	
doc:
	swag init