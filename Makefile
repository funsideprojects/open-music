
MESSAGING_PORT=5001

# Development

docker-up:
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml --env-file ./.env.development up -d

docker-down:
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml --env-file ./.env.development down

dev:
	gin -appPort 5000 -i -x bin run websocket/server.go

devm:
	gin -p 5001 -a 5011 -b bin/dev-messaging -t messaging -i --logPrefix messaging run messaging/server.go

build:
	rm -rf bin/dev-*
	go build -o bin/messaging messaging/server.go

start:
	go run server.go
