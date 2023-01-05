package main

import (
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

	strmInfo, err := js.AddStream(&nats.StreamConfig{
		Name:        "LOGS",
		Description: "watch for logs",
		Subjects:    []string{"LOGS.*"},
	})
	if err != nil {
		klog.Errorf(err.Error())
		return
	}

	consInfo, err := js.AddConsumer(strmInfo.Config.Name, &nats.ConsumerConfig{
		Durable:       "MONITOR",
		AckPolicy:     nats.AckExplicitPolicy,
		FilterSubject: "LOGS.V1",
	})
	if err != nil {
		klog.Errorf(err.Error())
		return
	}

	sub, err := js.PullSubscribe(consInfo.Config.FilterSubject, consInfo.Name, nats.BindStream(consInfo.Stream))
	if err != nil {
		klog.Errorf(err.Error())
		return
	}
	defer func() {
		err := sub.Unsubscribe()
		if err != nil {
			klog.Errorf(err.Error())
		}
	}()

	for {
		msgs, err := sub.Fetch(1)
		if err == nats.ErrTimeout {
			continue
		}
		if err != nil {
			klog.Errorf(err.Error())
			break
		}

		if len(msgs) == 0 {
			continue
		}

		err = msgs[0].Ack()
		if err != nil {
			klog.Errorf(err.Error())
			break
		}

		klog.Infof("Message received: %s ", string(msgs[0].Data))

	}
}
