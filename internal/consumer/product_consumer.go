package consumer

import (
	"context"
	"encoding/json"
	"gitee.com/phper95/pkg/es"
	"gitee.com/phper95/pkg/mq"
	"github.com/Shopify/sarama"
	"product-consumer/global"
	"strconv"
)

var productConsumer *mq.Consumer

func StartConsumer() {
	var err error
	productConsumer, err = mq.StartKafkaConsumer(global.CONFIG.Kafka.Hosts, []string{global.Topic},
		"product-consumer", nil, MsgHandler)
	if err != nil {
		global.LOG.Error("StartKafkaConsumer error", err)
		panic(err)
	}
}

func MsgHandler(msg *sarama.ConsumerMessage) (bool, error) {
	mq.KafkaStdLogger.Printf("partion: %d ; offset : %d; msg : %s",
		msg.Partition, msg.Offset, string(msg.Value))
	productMsg := ProductMsg{}
	err := json.Unmarshal(msg.Value, &productMsg)
	if err != nil {
		//格式异常的数据，回到队列也不会解析成功
		global.LOG.Error("Unmarshal error", err, string(msg.Value))
		return true, nil
	}
	mq.KafkaStdLogger.Printf("product: %+v", productMsg)
	productIndex := productMsg.ProductIndex
	esClient := es.GetClient(es.DefaultClient)
	switch productMsg.Operation {
	case global.OperationCreate, global.OperationOnSale:
		if productMsg.IsShow == 1 {
			esClient.BulkCreate(global.IndexName, strconv.FormatInt(productIndex.Id, 10),
				strconv.Itoa(productIndex.CateId), productIndex)
		}
	case global.OperationUpdate:
		if productMsg.IsShow == 0 {
			err := esClient.DeleteRefresh(context.Background(), global.IndexName,
				strconv.FormatInt(productIndex.Id, 10),
				strconv.Itoa(productIndex.CateId))
			if err != nil {
				global.LOG.Error("DeleteRefresh error", err, "id", productIndex.Id)
			}
		} else {
			esClient.BulkCreate(global.IndexName, strconv.FormatInt(productIndex.Id, 10),
				strconv.Itoa(productIndex.CateId), productIndex)
		}
	case global.OperationUnSale, global.OperationDelete:
		err := esClient.DeleteRefresh(context.Background(), global.IndexName,
			strconv.FormatInt(productIndex.Id, 10),
			strconv.Itoa(productIndex.CateId))
		if err != nil {
			global.LOG.Error("DeleteRefresh error", err, "id", productIndex.Id)
		}
	}

	return true, nil
}

func CloseProductConsumer() {
	if productConsumer != nil {
		productConsumer.Close()
	}
}
