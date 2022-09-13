package consumer

import (
	"encoding/json"
	"fmt"
	"gitee.com/phper95/pkg/es"
	"gitee.com/phper95/pkg/mq"
	"gitee.com/phper95/pkg/strutil"
	"github.com/Shopify/sarama"
	"mail-consumer/global"
)

var emailConsumer *mq.Consumer

func StartConsumer() {
	var err error
	emailConsumer, err = mq.StartKafkaConsumer(global.CONFIG.Kafka.Hosts, []string{global.Topic},
		"mail-consumer", nil, MsgHandler)
	if err != nil {
		global.LOG.Error("StartKafkaConsumer error", err)
		panic(err)
	}
}

func MsgHandler(msg *sarama.ConsumerMessage) (bool, error) {
	mq.KafkaStdLogger.Printf("partion: %d ; offset : %d; msg : %s",
		msg.Partition, msg.Offset, string(msg.Value))
	mailMsg := MailMsg{}
	err := json.Unmarshal(msg.Value, &mailMsg)
	if err != nil {
		//格式异常的数据，回到队列也不会解析成功
		global.LOG.Error("Unmarshal error", err, string(msg.Value))
		return true, nil
	}
	mq.KafkaStdLogger.Printf("mail msg: %+v", mailMsg)
	mailIndex := mailMsg.MailIndex
	esClient := es.GetClient(es.DefaultClient)
	routing := fmt.Sprintf("%d_%d", mailMsg.Uid, mailMsg.Id)
	if mailMsg.Operation == global.OperationDelete {
		esClient.BulkDelete(global.IndexName, routing, strutil.Int64ToString(mailMsg.Uid))
	} else {
		esClient.BulkCreate(global.IndexName, routing, strutil.Int64ToString(mailMsg.Uid), mailIndex)
	}
	return true, nil
}

func CloseConsumer() {
	if emailConsumer != nil {
		emailConsumer.Close()
	}
}
