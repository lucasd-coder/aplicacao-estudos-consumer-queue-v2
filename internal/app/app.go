package app

import (
	"sync"

	"github.com/lucasd-coder/aplicacao-estudos-consumer-queue-v2/config"
	"github.com/lucasd-coder/aplicacao-estudos-consumer-queue-v2/internal/processor"
	"github.com/lucasd-coder/aplicacao-estudos-consumer-queue-v2/internal/service"
	"github.com/lucasd-coder/aplicacao-estudos-consumer-queue-v2/internal/utils"
	"github.com/lucasd-coder/aplicacao-estudos-consumer-queue-v2/pkg/logger"
	"github.com/lucasd-coder/aplicacao-estudos-consumer-queue-v2/pkg/queuev2"

	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {
	// Log config
	logger.SetUpLog(cfg)

	// Http server
	engine := gin.New()
	engine.Use(gin.Recovery())

	// Routers
	handler := engine.Group("/" + cfg.Name)
	handler.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	})

	queuev2.SetUpSQS(cfg)

	notificationService := service.QueueConsumerService{}

	queueList := utils.QueueList(cfg.QueueList)

	go engine.Run(":" + cfg.Port)

	nQueues := len(queueList)
	wg := sync.WaitGroup{}
	wg.Add(nQueues)

	for i := range queueList {
		queues := utils.QueueNameUtils(cfg.Sqs.Prefix, "", queueList[i])

		go func() {
			defer wg.Done()
			consumer := processor.New(queues, func(msg types.Message) error {
				err := notificationService.QueueConsumer(&msg)
				return err
			}, &processor.Config{
				Receivers:                   cfg.Receivers,
				SqsMaxNumberOfMessages:      cfg.MaxNumberOfMessages,
				SqsMessageVisibilityTimeout: cfg.MessageVisibilityTimeout,
				PollDelayInMilliseconds:     cfg.PollDelayInMilliseconds,
			})

			consumer.Start()
		}()
	}

	wg.Wait()
}
