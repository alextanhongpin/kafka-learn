# Error Handling

- what is ACK and NACK?
- some libraries like golang's kafka pkg doesn't support NACK, and the recommendation is to create an `error topic` [^1]
- three main strategy for handling errors [^2]
  - fail fast (default): stops the application and marks it as unhealthy
  - ignore: continues the processing even if there is an error
  - dead-letter-topic: sends failed message to another topic for investigation. Name of the error topic can be `dead-letter-topic-$topic-name`

## References

- https://www.confluent.io/blog/error-handling-patterns-in-kafka/
- https://www.confluent.io/blog/kafka-connect-deep-dive-error-handling-dead-letter-queues/
- https://blogs.perficient.com/2021/02/15/kafka-consumer-error-handling-retry-and-recovery/

[^1]: https://github.com/segmentio/kafka-go/issues/575
[^2]: https://quarkus.io/blog/kafka-failure-strategy/
