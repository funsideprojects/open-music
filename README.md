# README

## Scripts

- Development

```bash
docker-compose -f docker-compose.yml -f docker-compose.dev.yml --env-file ./.env.development up -d
```

- Production

```bash
docker-compose -f docker-compose.yml -f docker-compose.production.yml --env-file ./.env.production up -d
```

## Docker Images

- [Vault](https://hub.docker.com/_/vault)
- [RabbitMQ](https://hub.docker.com/_/rabbitmq)
  - _15672_ is the default port for RabbitMQ GUI, _5672_ for RabbitMQ message broker.
  - [Environments vars](https://www.rabbitmq.com/configure.html)
- [Mongo DB](https://hub.docker.com/_/mongo)
- [Mongo Express](https://hub.docker.com/_/mongo-express)
- [Casssandra DB](https://hub.docker.com/_/cassandra)

## Useful libs

- [gin](github.com/codegangsta/gin): Live reload
- [websoket](github.com/gorilla/websocket)
- [web framework](github.com/labstack/echo)

## Useful links

- [Cassandra Docker Auth](https://hopding.com/cassandra-authentication-in-docker-container)
- [Docker-compose .env file](https://docs.docker.com/compose/environment-variables/)
- [Extends docker-compose files](https://docs.docker.com/compose/extends/)
