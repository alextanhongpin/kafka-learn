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

## Best practices

- always set `auto.commit.enable` to `false`. We need the `at least once` guaranty, not the `at most once`. Therefore we need to commit the offsets manually. 

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

## Kafka vs Message Queue

Kafka's model is __dumb broker, smart consumer__. Kafka is more suitable for publishing events (implementing Event Notification). However, the client needs to implement their own mechanism to ensure that the same message is not processed twice, e.g. by storing the last consumed offset, or by caching the id of the message that has been successfully processed.

Message queue's model is __smart broker, dumb consumer__. Message queue is more suitable for publishing commands, because we can acknowledge the message (commands) that are processed successfully, or retrying failed operations.

## Other tools

- kafkacat

## References

- https://semaphoreci.com/community/tutorials/writing-and-testing-an-event-sourcing-microservice-with-kafka-and-go
- https://medium.com/event-driven-utopia/understanding-kafka-topic-partitions-ae40f80552e8
- https://medium.com/fintechexplained/12-best-practices-for-using-kafka-in-your-architecture-a9d215e222e3
- https://aiven.io/blog/tips-for-designing-payloads
