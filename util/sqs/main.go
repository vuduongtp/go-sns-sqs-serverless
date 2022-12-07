package sqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

const (
	defaultRegion = "ap-southeast-1"
)

// New initializes SQS service with default config
func New() *Service {
	return &Service{
		sqs: sqs.New(session.New(&aws.Config{Region: aws.String(defaultRegion)})),
	}
}

// Service represents the snsutil service
type Service struct {
	sqs *sqs.SQS
}

// GetQueueURL by name
func (s *Service) GetQueueURL(name string) (*string, error) {
	urlResult, err := s.sqs.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &name,
	})
	if err != nil {
		return nil, err
	}

	return urlResult.QueueUrl, nil
}

// ReceiveMessage func
func (s *Service) ReceiveMessage(queueURL string, number int64) ([]*sqs.Message, error) {
	msgResult, err := s.sqs.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            aws.String(queueURL),
		MaxNumberOfMessages: aws.Int64(number),
		VisibilityTimeout:   aws.Int64(240),
	})

	if err != nil {
		return nil, err
	}

	return msgResult.Messages, nil
}

// DeleteMessage func
func (s *Service) DeleteMessage(queueURL, receiptHandle string) error {
	_, err := s.sqs.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueURL),
		ReceiptHandle: aws.String(receiptHandle),
	})

	if err != nil {
		return err
	}

	return nil
}
