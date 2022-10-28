package processor

import (
	"context"
	"time"

	"github.com/lucasd-coder/aplicacao-estudos-consumer-queue-v2/pkg/logger"
	"github.com/lucasd-coder/aplicacao-estudos-consumer-queue-v2/pkg/queuev2"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

// SqsReceiver defines the struct that polls messages from AWS SQS
type SqsReceiver struct {
	queueURL                string
	messagesChannel         chan []types.Message
	visibilityTimeout       int64
	maxNumberOfMessages     int64
	pollDelayInMilliseconds int
}

func (r *SqsReceiver) applyBackPressure() {
	time.Sleep(time.Millisecond * time.Duration(r.pollDelayInMilliseconds))
}

func (r *SqsReceiver) receiveMessages() {
	queue := queuev2.GetClient()
	retry := time.Duration(30) * time.Second
	for {
		result, err := queue.ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
			AttributeNames: []types.QueueAttributeName{
				"SentTimestamp",
			},
			MessageAttributeNames: []string{
				string(types.QueueAttributeNameAll),
			},
			QueueUrl:            &r.queueURL,
			MaxNumberOfMessages: 1,
			VisibilityTimeout:   int32(r.visibilityTimeout),
		})
		if err != nil {
			logger.Log.Error("Could not read from queue ", err)
			time.Sleep(retry)
			continue
		}

		if len(result.Messages) > 0 {
			messages := result.Messages
			r.messagesChannel <- messages
		}

		r.applyBackPressure()
	}
}
