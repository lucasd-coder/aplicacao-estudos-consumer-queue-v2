package service

import (
	"github.com/lucasd-coder/aplicacao-estudos-consumer-queue-v2/pkg/logger"

	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type QueueConsumerService struct{}

func NewQueueConsumerService() *QueueConsumerService {
	return &QueueConsumerService{}
}

func (n *QueueConsumerService) QueueConsumer(msg *types.Message) error {
	logger.Log.Info("Message SQSBody:", *msg.Body)

	return nil
}
