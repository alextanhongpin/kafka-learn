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

## References

https://semaphoreci.com/community/tutorials/writing-and-testing-an-event-sourcing-microservice-with-kafka-and-go
