# Project rabbitmq-blueprint

Concept for implementing RabbitMQ in the go-blueprint as an advanced flag.

```bash
docker compose -f docker-compose-rabbitmq.yml up --build

docker compose -f docker-compose-rabbitmq.yml down --volumes
```

Publisher
```bash
curl -X POST -d 'message=test' localhost:8088/publish
```

Consumer
```bash
docker logs -f consumer-1
```

RannitMQ UI
```bash
localhost:15672
```


## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

Create containers
```bash
make docker-run
```

Shutdown containers
```bash
make docker-down
```

live reload the application
```bash
make watch
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```