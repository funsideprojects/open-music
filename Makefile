
MESSAGING_PORT=5001

dev:
	gin -appPort 5000 -i -x bin run websocket/server.go

devm:
	gin -p 5001 -a 5011 -b bin/dev-messaging -t messaging -i --logPrefix messaging run messaging/server.go

build:
	go build -o bin/messaging messaging/server.go

start:
	go run server.go
