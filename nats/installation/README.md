### Install
* NATS cli: https://github.com/nats-io/natscli


`go install github.com/nats-io/natscli/nats@latest` <br>
* for specific version: `go install github.com/nats-io/natscli/nats@v0.0.33`

## Commands
### Core NATS
* `nats server run` # to run the nats server
* `nats context select nats_development` # to select context account
* `nats sub <subject_name>` # to create a subscriber. The subscriber will only listen for message from publisher.
* `nats pub <subject_name> "message"` # to create a publisher. For publishing something it must have some subscriber to listen that subject
* `nats req <subject.name> "message"` # sends a message
* `nats reply <subject.name> "message"` # listen and sends the `message` to the requester.
* 
