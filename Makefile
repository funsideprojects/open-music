build:
	go build -o bin/main server.go

start:
	go run server.go

dev:
	gin -p 3000 run server.go