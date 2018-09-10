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
var c chan string

func main() {
	c = make(chan string, 20000)

	iLimit := 1
	iLimit_sub := 10000 * 800

	t1 := time.Now()
	for i := 0; i < iLimit; i++ {
		wg.Add(1)
		go run(i, iLimit_sub)
	}

	//-send-----------------

	for i := 0; i < iLimit*iLimit_sub; i++ {
		info := fmt.Sprintf("id: %d", i)
		c <- info
	}
	fmt.Println("------after setData-------------------")
	wg.Wait()
	fmt.Println("#### count:", iLimit*iLimit_sub)
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

	for {
		ok := do_send(producer, topic, routine_id)
		if !ok {
			time.Sleep(time.Millisecond * 10)
		}
	} //for
	//		k = partition + offset
	fmt.Println("finished-->sub:", routine_id)
}

func get_messages(topic *string, c <-chan string) (lstMsg []*sarama.ProducerMessage) {
	//	fmt.Println("### enter get_messages ###")
	lstMsg = []*sarama.ProducerMessage{}
	timeout := make(chan bool)

	for i := 0; i < 10000; i++ {
		go wait(timeout, 100)
		select {
		case s := <-c:
			msg := &sarama.ProducerMessage{
				Topic: *topic,
				Value: sarama.StringEncoder(s)}
			lstMsg = append(lstMsg, msg)
		case <-timeout:
			goto OUTER
		} //select
	} //for
OUTER:
	if len(lstMsg) > 0 {
		//	fmt.Println("get_message is return--count:", len(lstMsg))
		return lstMsg
	}

	//	fmt.Println("get _message 00000000000")
	return nil
}

func wait(timeout chan<- bool, ms time.Duration) {
	time.Sleep(time.Millisecond * ms)
	timeout <- true
}

func do_send(producer sarama.SyncProducer, topic *string, routine_id int) bool {
	lstMsg := get_messages(topic, c)
	if lstMsg == nil {
		return false
	}

	if len(lstMsg) > 0 {
		err := producer.SendMessages(lstMsg)
		if err != nil {
			fmt.Println("###############################sendError######")
			panic(err)
		}
		fmt.Println("send queen: ", routine_id, "count:", len(lstMsg))
	}

	return true
}
