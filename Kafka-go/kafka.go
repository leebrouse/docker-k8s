package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"
)

var (
	reader   *kafka.Reader
	topic    = "user_click"
	maxtimes = 3
)

// producer
func writeKafka(ctx context.Context) {
	writer := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  topic,
		Balancer:               &kafka.Hash{},
		WriteTimeout:           1 * time.Second,
		RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: true,
	}

	defer writer.Close()

	for i := 0; i < maxtimes; i++ {
		if err := writer.WriteMessages(
			ctx,
			kafka.Message{Key: []byte("1"), Value: []byte("Hello World!")},
			kafka.Message{Key: []byte("2"), Value: []byte("bye World!")},
			kafka.Message{Key: []byte("3"), Value: []byte("new hello")},
			kafka.Message{Key: []byte("4"), Value: []byte("Lee brouse")},
			kafka.Message{Key: []byte("5"), Value: []byte("docker k8s")},
		); err != nil {
			if err == kafka.KafkaStorageError {
				time.Sleep(2 * time.Second)
				continue
			} else {
				log.Fatal("write message error: %v", err)
			}
		} else {
			break
		}
	}

}

// consumer
func readKafka(ctx context.Context) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{"localhost:9092"},
		Topic:          topic,
		CommitInterval: 1 * time.Second,
		GroupID:        "rec_team",
		StartOffset:    kafka.FirstOffset,
	})

	defer reader.Close()

	for {
		if msg, err := reader.ReadMessage(ctx); err != nil {
			log.Fatal("read message error: %v", err)
			break
		} else {
			fmt.Printf("topic=%s,partition=%d,offset=%d,key=%s,value=%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
		}
	}
}

// 需要监听信息2和15，当收到信号时关闭reader
func listenSignal() {
	c :=make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	sig:= <-c
	fmt.Printf("Got signal:%s", sig.String())
	if reader!=nil{
		reader.Close()
	}

	os.Exit(0)
}

func main() {
	ctx := context.Background()
	writeKafka(ctx)

	// go listenSignal()
	// readKafka(ctx)

}
