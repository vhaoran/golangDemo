package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	//
	iLimit := 60
	iLimit_sub := 10000 * 5
	t1 := time.Now()

	for i := 0; i < iLimit; i++ {
		wg.Add(1)
		go run(i, iLimit_sub)
	}
	wg.Wait()
	fmt.Println("seconds:", (time.Now().Unix() - t1.Unix()))
}

func run(routine_id int, iLimit int) {
	defer wg.Add(-1)
	var (
		brokerList = kingpin.Flag("brokerList", "List of brokers to connect").Default("localhost:9092").Strings()
		topic      = kingpin.Flag("topic", "Topic name").Default("test").String()
		maxRetry   = kingpin.Flag("maxRetry", "Retry limit").Default("5").Int()
	)

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
	for i := 0; i < iLimit; i++ {
		info := fmt.Sprintf("id: %d message send %d", routine_id, i)
		msg := &sarama.ProducerMessage{
			Topic: *topic,
			Value: sarama.StringEncoder(info),
		}
		//partition, offset, err := producer.SendMessage(msg)
		_, _, err := producer.SendMessage(msg)
		if err != nil {
			panic(err)
		}

		//		k = partition + offset
	}
	//	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", *topic, partition, offset)

	fmt.Println("-------time--", time.Now(),
		"\n\rms:", (time.Now().Unix() - t1.Unix()), "  loop:", iLimit)
}
