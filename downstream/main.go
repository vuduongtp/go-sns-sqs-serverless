package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/vuduongtp/go-sns-sqs-serverless/util/print"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// SQSMessageData model
type SQSMessageData struct {
	Message          string `json:"Message"`
	MessageID        string `json:"MessageId"`
	Signature        string `json:"Signature"`
	SignatureVersion string `json:"SignatureVersion"`
	SigningCertURL   string `json:"SigningCertURL"`
	Timestamp        string `json:"Timestamp"`
	TopicArn         string `json:"TopicArn"`
	Type             string `json:"Type"`
	UnsubscribeURL   string `json:"UnsubscribeURL"`
}

func main() {
	lambda.Start(func(ctx context.Context, sqsEvent events.SQSEvent) (string, error) {
		err := Run(ctx, sqsEvent)
		if err != nil {
			return "ERROR", fmt.Errorf("ERROR: %+v", err)
		}

		return "OK", nil
	})
}

// Run Lambda
func Run(ctx context.Context, sqsEvent events.SQSEvent) error {
	print.PrettyPrint(sqsEvent)
	for _, mess := range sqsEvent.Records {
		sqsRes := SQSMessageData{}
		if err := json.Unmarshal([]byte(mess.Body), &sqsRes); err != nil {
			return err
		}

		intf := make(map[string]interface{})
		if err := json.Unmarshal([]byte(sqsRes.Message), &intf); err != nil {
			return err
		}

		print.PrettyPrint(intf)
	}

	return nil
}
