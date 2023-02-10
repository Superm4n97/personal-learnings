### Run the server
first run the server
```bash
go run server-nats.go
```

### Run the publisher
You can send any string in this publisher, the message will be sent in `LOGS.V1` stream.
```bash
go run jstrm-publisher.go
```

### Run the consumer
The subscriber will read message from the `LOGS.V1` stream.
```bash
go run jstrm-consumer.go
```