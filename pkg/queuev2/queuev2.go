package queuev2

import (
	"context"
	"time"

	aplicationConfig "github.com/lucasd-coder/aplicacao-estudos-consumer-queue-v2/config"
	"github.com/lucasd-coder/aplicacao-estudos-consumer-queue-v2/pkg/logger"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

var client *sqs.Client

func SetUpSQS(aplicationConfig *aplicationConfig.Config) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(aplicationConfig.SqsRegion),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				PartitionID:       "aws",
				URL:               aplicationConfig.SqsHost,
				SigningRegion:     aplicationConfig.SqsRegion,
				HostnameImmutable: true,
			}, nil
		})))
	if err != nil {
		logger.Log.Error(err)
	}

	client = sqs.NewFromConfig(cfg)
}

func GetClient() *sqs.Client {
	return client
}
