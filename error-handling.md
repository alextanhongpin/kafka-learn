# Error Handling

- what is ACK and NACK?
- some libraries like golang's kafka pkg doesn't support NACK, and the recommendation is to create an `error topic` [^1]

## References

- https://quarkus.io/blog/kafka-failure-strategy/
- https://www.confluent.io/blog/error-handling-patterns-in-kafka/
- https://www.confluent.io/blog/kafka-connect-deep-dive-error-handling-dead-letter-queues/
- https://blogs.perficient.com/2021/02/15/kafka-consumer-error-handling-retry-and-recovery/

[^1]: https://github.com/segmentio/kafka-go/issues/575
