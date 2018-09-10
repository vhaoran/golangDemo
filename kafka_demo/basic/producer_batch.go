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
	k := 0
	iLimit := 100000 * 5
	lstMsg := []*sarama.ProducerMessage{}
	for i := 0; i < iLimit; i++ {
		info := fmt.Sprintf("message send %d", i)
		msg := &sarama.ProducerMessage{
			Topic: *topic,
			Value: sarama.StringEncoder(info),
		}
		//partition, offset, err := producer.SendMessage(msg)
		lstMsg = append(lstMsg, msg)
		k++
	}

	err = producer.SendMessages(lstMsg)
	if err != nil {
		panic(err)
	}

	//		k = partition + offset

	//	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", *topic, partition, offset)

	fmt.Println("-------time--", time.Now(),
		"\n\rms:", (time.Now().Unix() - t1.Unix()), "  loop:", iLimit)

}
