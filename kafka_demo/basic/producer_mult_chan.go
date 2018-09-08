package main

//
//在kafaka连接资源有限的情况下，通过有限的go roution发送消息
// 缺点：有缓冲区，可能在机器发现问题时消息未发出，
// 优点：不需要等等，直接发送
//
import (
	"flag"
	"fmt"
	"github.com/Shopify/sarama"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	//	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	//	if len(os.Args) < 2 {
	//		fmt.Println("args is p1 and p2")
	fmt.Println("----args----", flag.Args(), "###  len:", len(flag.Args()))
	//		return
	//	}

	iLimit := flag.Int("c", 5, "p1为goroutine count")
	iLimit_sub := flag.Int("n", 10000*5, "线程发送消息数量")
	flag.Parse()
	fmt.Println("----args----", flag.Args(), "###  len:", len(flag.Args()))

	t1 := time.Now()
	for i := 0; i < *iLimit; i++ {
		wg.Add(1)
		go run(i, *iLimit_sub)
	}
	wg.Wait()
	fmt.Println("#### count:", (*iLimit)*(*iLimit_sub))
	fmt.Println("all ######seconds:", (time.Now().Unix() - t1.Unix()))
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

	for i := 0; i < iLimit; i++ {
		t1 := time.Now()
		info := fmt.Sprintf("id: %d message send %d time: %d-%d-%d %d:%d:%d", routine_id, i, t1.Year(), t1.Month(), t1.Day(), t1.Hour(), t1.Minute(), t1.Second())
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

	fmt.Println("finished-->sub:", routine_id, "--second:", (time.Now().Unix() - t1.Unix()))

}
