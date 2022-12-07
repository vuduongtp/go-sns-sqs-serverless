package sns

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

const (
	defaultRegion = "ap-southeast-1"
)

// MessageSNS to send to topic
type MessageSNS struct {
	Default string `json:"default"`
}

// New initializes SNS service with default config
func New() *Service {
	return &Service{
		sns: sns.New(session.New(&aws.Config{Region: aws.String(defaultRegion)})),
	}
}

// Service represents the snsutil service
type Service struct {
	sns *sns.SNS
}

// SendMsgToTopic sends message to a topic. The "msg" must be in json
func (s *Service) SendMsgToTopic(topicARN string, msg interface{}) (string, error) {
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	message := MessageSNS{
		Default: string(jsonMsg),
	}
	messageBytes, _ := json.Marshal(message)
	messageStr := string(messageBytes)
	fmt.Println(messageStr)
	output, err := s.sns.Publish(&sns.PublishInput{
		Message:          aws.String(messageStr),
		MessageStructure: aws.String("json"),
		TopicArn:         aws.String(topicARN),
	})

	if err != nil {
		return "", err
	}

	return *output.MessageId, nil
}
