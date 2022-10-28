package processor

import (
	"time"

	"github.com/lucasd-coder/aplicacao-estudos-consumer-queue-v2/pkg/logger"
	"github.com/lucasd-coder/aplicacao-estudos-consumer-queue-v2/pkg/queuev2"

	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

// Consumer holds the consumer data
type Consumer struct {
	queueURL        string
	messagesChannel chan []types.Message
	handler         func(m types.Message) error
	config          *Config
	receiver        SqsReceiver
}

// Config holds the configuration for consuming and processing the queue
type Config struct {
	SqsMaxNumberOfMessages      int64
	SqsMessageVisibilityTimeout int64
	Receivers                   int
	PollDelayInMilliseconds     int
}

// New creates a new Queue consumer
func New(queueURL string, handler func(m types.Message) error, config *Config) Consumer {
	c := make(chan []types.Message)

	r := SqsReceiver{
		queueURL:                queueURL,
		messagesChannel:         c,
		visibilityTimeout:       config.SqsMessageVisibilityTimeout,
		maxNumberOfMessages:     config.SqsMaxNumberOfMessages,
		pollDelayInMilliseconds: config.PollDelayInMilliseconds,
	}

	return Consumer{
		queueURL:        queueURL,
		messagesChannel: c,
		handler:         handler,
		config:          config,
		receiver:        r,
	}
}

// Start initiates the queue consumption process
func (c *Consumer) Start() {
	time.Sleep(time.Second * 2)
	logger.Log.Info("Starting to consume: ", c.queueURL)
	c.startReceivers()
	c.startProcessor()
}

// startReceivers starts N (defined in NumberOfMessageReceivers) goroutines to poll messages from SQS
func (c *Consumer) startReceivers() {
	for i := 0; i < c.config.Receivers; i++ {
		go c.receiver.receiveMessages()
	}
}

// startProcessor starts a goroutine to handle each message from messagesChannel
func (c *Consumer) startProcessor() {
	queue := queuev2.GetClient()

	p := Processor{
		queueURL: c.queueURL,
		queue:    queue,
		handler:  c.handler,
	}

	for messages := range c.messagesChannel {
		go p.processMessages(messages)
	}
}

// SetPollDelay increases time between message poll
func (c *Consumer) SetPollDelay(delayBetweenPoolsInMilliseconds int) {
	c.receiver.pollDelayInMilliseconds = delayBetweenPoolsInMilliseconds
}
