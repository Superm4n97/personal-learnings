package main

import (
	"github.com/nats-io/nats.go"
	"k8s.io/klog/v2"
)

const publisherSubjectName = "logs"

func main() {

	// connect to nats server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		klog.Errorf(err.Error())
	}

	// simple publisher
	err = nc.Publish(publisherSubjectName, []byte("hello world"))
	if err != nil {
		klog.Errorf(err.Error())
	}
	klog.Infof("exiting publisher...")
}
