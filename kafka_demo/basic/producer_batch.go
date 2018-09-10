package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"time"
)

var (
	brokerList = kingpin.Flag("brokerList", "List of brokers to connect").Default("localhost:9092").Strings()
	topic      = kingpin.Flag("topic", "Topic name").Default("test").String()
	maxRetry   = kingpin.Flag("maxRetry", "Retry limit").Default("5").Int()
)

func main() {
	kingpin.Parse()
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = *maxRetry
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(*brokerList, config)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			panic(err)
		}
	}()

	t1 := time.Now()
	fmt.Println("------------begin---------time--", t1)
	k := 0              //counter
	h := 1000           //outer llops
	iLimit := 10000 * 1 //send each times

	for j := 0; j < h; j++ {
		lstMsg := []*sarama.ProducerMessage{}
		for i := 0; i < iLimit; i++ {
			info := fmt.Sprintf("message send %d", i)
			msg := &sarama.ProducerMessage{
				Topic: *topic,
				Value: sarama.StringEncoder(info),
			}
			lstMsg = append(lstMsg, msg)
			k++
		} //first for

		err = producer.SendMessages(lstMsg)
		if err != nil {
			panic(err)
		}
	}
	//not to here
	fmt.Println("ms:", (time.Now().Unix() - t1.Unix()), "  loop:", iLimit, " all count:", k)

}
