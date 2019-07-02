package mqtt

import (
	"flag"
	"fmt"
	"log"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type subscribe struct {
	cli *client

	topic string
	qos   uint
	idgen func() string
	debug bool
}

func NewSubscribeCommand(cli *client, args []string) Command {
	c := &subscribe{}
	fs := flag.NewFlagSet("subscribe", flag.ExitOnError)
	setOpt(fs, &cli.opt)
	fs.StringVar(&c.topic, "topic", "fperf-topic", "Topic to subscribe")
	fs.UintVar(&c.qos, "qos", 1, "QoS should be 0, 1, 2")
	fs.BoolVar(&c.debug, "debug", false, "Print interactive information")
	if err := fs.Parse(args[1:]); err != nil {
		log.Fatal(err)
	}
	c.cli = cli
	c.idgen = idgenerator()
	return c
}

func (c *subscribe) Exec() error {
	qos := c.qos
	topic := c.topic
	cli := c.cli

	handler := func(client MQTT.Client, msg MQTT.Message) {
		if c.debug {
			fmt.Printf("recv msg = %s\n", string(msg.Payload()))
		}
	}
	mq := cli.cli
	idstr := c.idgen()
	if token := mq.Subscribe(topic+"-"+idstr, byte(qos), handler); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
