package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"k8s.io/klog/v2"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		klog.Errorf(err.Error())
		return
	}

	js, err := nc.JetStream()
	if err != nil {
		klog.Errorf(err.Error())
		return
	}

	_, err = js.AddStream(&nats.StreamConfig{
		Name:        "LOGS",
		Description: "watch for logs",
		Subjects:    []string{"LOGS.*"},
	})
	if err != nil {
		klog.Errorf(err.Error())
		return
	}

	for {
		var msg string
		fmt.Scanln(&msg)

		if msg == "exit" {
			break
		}

		natsMsg := &nats.Msg{
			Subject: "LOGS.V1",
			Data:    []byte(msg),
		}
		_, err := js.PublishMsg(natsMsg)
		if err != nil {
			klog.Errorf(err.Error())
			break
		}
	}
}
