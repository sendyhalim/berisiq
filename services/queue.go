package services

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type queueClient struct {
	sqsClient *sqs.SQS
}

var Queue *queueClient

func (q *queueClient) FilterQueuesByPrefix(prefix string) ([]string, error) {
	output, err := q.sqsClient.ListQueues(&sqs.ListQueuesInput{
		QueueNamePrefix: &prefix,
	})

	if err != nil {
		return nil, err
	}

	s := make([]string, len(output.QueueUrls))

	for index, value := range output.QueueUrls {
		s[index] = *value
	}

	return s, nil
}

func (q *queueClient) SendMessage(queueName string, message string) error {
	_, err := q.sqsClient.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    &queueName,
		MessageBody: &message,
	})

	return err
}

func init() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"",
		),
	})

	if err != nil {
		log.Fatalf("Could not initiate aws session %s", err)
	}

	Queue = &queueClient{
		sqsClient: sqs.New(sess),
	}
}
