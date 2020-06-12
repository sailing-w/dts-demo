package main

import (
	"bytes"
	"fmt"
	"github.com/sanity-io/litter"
	"io"
	"os"
	"os/signal"
	"time"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

var (
	r io.Reader

	// 仅需改动以下配置即可
	// ***********************************************************
	kafkaUser       = "data_brigde_sync"
	kafkaPassWord   = "Medlinker123"
	kafkaTopic      = "topic"
	kafkaGroupId    = "groupid"
	kafkaBrokerList = []string{"localhost:9020"}
	// ***********************************************************
)

func main() {
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Net.MaxOpenRequests = 100
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Net.SASL.Enable = true
	config.Net.SASL.User = fmt.Sprintf("%s-%s",
		"data_brigde_sync", "dtsb2tf52en28mlbel")
	config.Net.SASL.Password = "Medlinker123"
	config.Version = sarama.V0_11_0_0

	consumer, err := cluster.NewConsumer([]string{"dts-cn-hangzhou.aliyuncs.com:18001"},
		"dtsb2tf52en28mlbel", []string{"cn_hangzhou_i_bp1ijo6mcpzxkvr3h0in_3306_meddev"}, config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()
	// trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// consume errors
	go func() {
		for err := range consumer.Errors() {
			panic(err)
		}
	}()

	// consume notifications
	go func() {
		for ntf := range consumer.Notifications() {
			fmt.Println("Rebalanced: ", litter.Sdump(ntf))
		}
	}()

	// consume messages, watch signals
	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				r = bytes.NewReader(msg.Value)
				fmt.Printf("key: %s, value: %s\n", string(msg.Key), string(msg.Value))

				//record, err := avro.DeserializeRecord(r)
				//if err != nil {
				//	log.Fatal("avro.DeserializeRecord is err: ", err)
				//}
				//if record.Operation == avro.OperationUPDATE {
				//	fmt.Println(litter.Sdump(record.Tags))
				//}
				//t := avro.NewRecord()
				//codec, err := goavro.NewCodec(t.Schema())
				//if err != nil {
				//	log.Fatal(err)
				//}
				//native, _, err := codec.NativeFromBinary(msg.Value)
				//if err != nil {
				//	log.Fatal("codec.NativeFromBinary error:", err)
				//}
				////fmt.Println("native:", native)
				//
				//_, err = codec.TextualFromNative(nil, native)
				//if err != nil {
				//	log.Fatal("codec.TextualFromNative: ", err)
				//}
				////fmt.Println("texual:", string(texual))

			}
		case <-signals:
			return
		}
	}
}
