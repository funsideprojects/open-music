# Development

docker-devup:
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml --env-file ./.env.development up -d

docker-devdown:
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml --env-file ./.env.development down

# Identity
build-identity:
	go build -o identity/.air/main identity/server.go
dev-identity:
	air -c identity/.air.toml

devm:
	gin -p 5001 -a 5011 -b bin/dev-messaging -t messaging -i --logPrefix messaging run messaging/server.go

build:
	rm -rf bin/dev-*
	go build -o bin/messaging messaging/server.go

start:
	go run server.go
