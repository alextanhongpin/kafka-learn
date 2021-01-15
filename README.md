## go-kafka

Example of running kafka with golang.

- [x] setup docker kafka single instance
- [x] running example with golang
- [ ] example with nodejs/python etc
- [ ] clustering
- [ ] in-depth kafka usage and scaling, consumer group etc...
- [ ] comparison against NATS, RabbitMQ
- [ ] production-scale integration

```bash
$ go get gopkg.in/Shopify/sarama.v1
```

Start docker:
```bash
$ docker-compose up -d
```

Test the connection:
```bash
$ nc -vz localhost:9092
```

## Connecting to Kafka from non-docker

Configure the following environment variables for kafka (e.g. docker-compose):
```yml
    environment:
      - KAFKA_CFG_ZOOKEEPER_CONNECT=zookeeper:2181
      - ALLOW_PLAINTEXT_LISTENER=yes
      - KAFKA_LISTENER_SECURITY_PROTOCOL_MAP=CLIENT:PLAINTEXT,EXTERNAL:PLAINTEXT
      - KAFKA_LISTENERS=CLIENT://:9092,EXTERNAL://:9093
      - KAFKA_ADVERTISED_LISTENERS=CLIENT://kafka:9092,EXTERNAL://localhost:9093
      - KAFKA_INTER_BROKER_LISTENER_NAME=CLIENT
```

Reference to the [configuration](https://docs.confluent.io/platform/current/kafka/multi-node.html).

## Other tools

- kafkacat

## References

https://semaphoreci.com/community/tutorials/writing-and-testing-an-event-sourcing-microservice-with-kafka-and-go
