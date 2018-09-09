package main

//
//在kafaka连接资源有限的情况下，通过有限的go roution发送消息
// 缺点：有缓冲区，可能在机器发现问题时消息未发出，
// 优点：不需要等等，直接发送
//
import (
	//	"flag"
	"fmt"
	"github.com/Shopify/sarama"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	//	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup
var c chan *string

func main() {
	c = make(chan *string, 100000)

	iLimit := 5
	iLimit_sub := 10000 * 5

	t1 := time.Now()
	for i := 0; i < iLimit; i++ {
		wg.Add(1)
		go run(i, iLimit_sub)
	}

	//-send-----------------
	for i := 0; i < iLimit*iLimit_sub; i++ {
		//减少等待
		go func(id int) {
			info := fmt.Sprintf("data_id: %d message send!", id)
			//push to chan
			c <- &info
		}(i)
	}
	fmt.Println("#### count:", iLimit*iLimit_sub)
	fmt.Println("all ######seconds:", (time.Now().Unix() - t1.Unix()))

	wg.Wait()
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

	//partition, offset, err := producer.SendMessage(ms
	//		for str := range c {
	total := 0
	for {
		//		fmt.Println("sub:", routine_id, "----before selecct")
		timeout := make(chan bool, 1)
		go func() {
			time.Sleep(time.Second * 10) // sleep one second
			timeout <- true
		}()

		skip := false
		select {
		case str, ok := <-c:
			if !ok {
				break
			}
			total += 1
			if total%1000 == 0 {
				fmt.Println("sbu: ", routine_id, "---total:", total)
			}
			//			fmt.Println("sub:", routine_id, "  total:", total)
			info := fmt.Sprintf("subid:%d c len:%d %s", routine_id, len(c), *str)
			msg := &sarama.ProducerMessage{
				Topic: *topic,
				Value: sarama.StringEncoder(info),
			}

			_, _, err := producer.SendMessage(msg)
			if err != nil {
				panic(err)
			}
		case <-timeout:
			fmt.Println("sub:", routine_id, "---time out")
			skip = true
			break
		} //select
		if skip {
			break
		}
	} //for

	//		k = partition + offset
	fmt.Println("finished-->sub:", routine_id, "--second:", (time.Now().Unix() - t1.Unix()), "---all count:----", total)
}
