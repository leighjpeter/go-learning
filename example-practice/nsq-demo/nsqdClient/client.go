package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/nsqio/go-nsq"
)

func main() {
	adds := []string{"127.0.0.1:32773"}
	config := nsq.NewConfig()
	config.MaxInFlight = 1000
	config.MaxBackoffDuration = 5 * time.Second
	config.DialTimeout = 10 * time.Second

	topicName := "testTopic2"
	c, _ := nsq.NewConsumer(topicName, "ch1", config)
	testHandler := &MyTestHandler{consumer: c}

	c.AddHandler(testHandler)
	/* if err := c.ConnectToNSQDs(adds); err != nil {
		panic(err)
	}
	stats := c.Stats()
	if stats.Connections == 0 {
		panic("stats report 0 connecttions (should be > 0)")
	}
	*/

	if err := c.ConnectToNSQLookupds(adds); err != nil {
		panic(err)
	}

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	fmt.Println("server is runing...")
	<-stop
}

type MyTestHandler struct {
	consumer *nsq.Consumer
}

func (my MyTestHandler) HandleMessage(message *nsq.Message) error {
	fmt.Println(string(message.Body))
	return nil
}
