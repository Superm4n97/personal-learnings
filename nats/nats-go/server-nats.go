package main

import (
	"github.com/nats-io/nats-server/v2/server"
	"k8s.io/klog/v2"
	"time"
)

func main() {
	natsServer, err := server.NewServer(&server.Options{
		JetStream: true,
	})
	if err != nil {
		klog.Errorf("failed to create NATS server")
		return
	}

	go natsServer.Start()
	defer natsServer.Shutdown()

	if !natsServer.ReadyForConnections(10 * time.Second) {
		klog.Errorf("can not start server")
		return
	}
	klog.Infof("NATS server is running...")
	select {}
}
