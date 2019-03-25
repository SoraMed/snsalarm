package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const webhook = "https://hooks.slack.com/services/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"

// LambdaAlarm handles an SNS Alert
func LambdaAlarm(ctx context.Context, snsEvent events.SNSEvent) {
	for _, e := range snsEvent.Records {
		msg := fmt.Sprintf("%v\n%v", e.SNS.Subject, e.SNS.Message)
		sendSlackMsg(msg)
	}
}

func sendSlackMsg(txt string) error {
	p := struct {
		Text string `json:"text"`
	}{Text: txt}
	js, err := json.Marshal(&p)
	if err != nil {
		return err
	}
	_, err = http.Post(webhook, "application/json", strings.NewReader(string(js)))
	return err
}

func main() {
	lambda.Start(LambdaAlarm)
}

/*
fmt.Printf("--- Event %v ---\n", i)
fmt.Printf("EventVersion: %v\n", e.EventVersion)
fmt.Printf("EventSubscriptionArn: %v\n", e.EventSubscriptionArn)
fmt.Printf("EventSource: %v\n", e.EventSource)
fmt.Printf("Signature: %v\n", e.SNS.Signature)
fmt.Printf("MessageID: %v\n", e.SNS.MessageID)
fmt.Printf("Type: %v\n", e.SNS.Type)
fmt.Printf("TopicArn: %v\n", e.SNS.TopicArn)
fmt.Printf("MessageAttributes: %v\n", e.SNS.MessageAttributes)
fmt.Printf("SignatureVersion: %v\n", e.SNS.SignatureVersion)
fmt.Printf("Timestamp: %v\n", e.SNS.Timestamp)
fmt.Printf("SigningCertURL: %v\n", e.SNS.SigningCertURL)
fmt.Printf("Message: %v\n", e.SNS.Message)
fmt.Printf("UnsubscribeURL: %v\n", e.SNS.UnsubscribeURL)
fmt.Printf("Subject: %v\n", e.SNS.Subject)
*/
