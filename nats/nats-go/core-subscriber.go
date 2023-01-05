package main

import (
	"github.com/nats-io/nats.go"
	"k8s.io/klog/v2"
)

const subscriberSubjectName = "logs"

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		klog.Errorf("failed to connect with nats server")
		return
	}

	nc.Subscribe(subscriberSubjectName, func(msg *nats.Msg) {
		klog.Infof("Message received: %s", string(msg.Data))
	})

	//// Responding to a request message
	//nc.Subscribe("requesting", func(msg *nats.Msg) {
	//	msg.Respond([]byte("hello"))
	//})

	klog.Infof("exiting subscriber...")
}
