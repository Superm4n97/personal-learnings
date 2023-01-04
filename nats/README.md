### Resource:
* Publisher Subscriber Model: https://www.youtube.com/watch?v=FMhbR_kQeHw
* Appscode public doc: https://docs.google.com/presentation/d/1kRuA_DD3B9-swdIXlEBz4kLfZ0dO11gh3zlfvJNSmns/edit#slide=id.p
* Official website: https://nats.io/
* CLI installation: https://github.com/nats-io/natscli
* Beginner tutorial: https://dev.to/karanpratapsingh/introduction-to-nats-cli-33nk
* Jetstream: https://docs.nats.io/nats-concepts/jetstream#streaming-temporal-decoupling-between-the-publishers-and-subscribers

## Notes:
### Core NATS
* NATS is a single technology that enables secure communication. NATS has `Publisher` and `Subscriber` scheme. Publisher publish a message to a `Subject` and the Subscriber has to subscribe to that Subject to receive the message.
* If Publisher sends message when the Subscriber is down, then the subscriber will miss the message, it is called `at-most once delivery`
* `Subjects` are identified by chars a-z,A-Z,0-9 and dot(.) separated.
* `Token`'s are dot separated component. and `*` matches a single token.
* A `Subscriber` can subscribe to multiple subjects using `wildcards`
* You can configure fan in/out style configuration, single subscriber to multiple publisher or multiple subscriber to single publisher.

### Jetstream (nats streaming)
* Streams are "message stores" (memory or file)
* Subscribers are called consumers and consume message from streams.

### NATS Go
* inport nats client: `github.com/nats-io/nats.go`
* 