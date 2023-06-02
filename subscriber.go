package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func main() {
	// Create a new AWS session
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-south-1"),                                                    // Replace with your desired AWS region
		Credentials: credentials.NewStaticCredentials("AWS ACCESS KEY", "AWS ACCESS KEY ID", ""), // Uncomment if not using shared credentials file
	})
	if err != nil {
		fmt.Println("Failed to create session:", err)
		return
	}

	// Create an SNS client
	svc := sns.New(sess)

	// Specify the topic ARN you want to subscribe to
	topicARN := "arn:aws:sns:ap-south-1:357044685191:Testing_SNS"

	// Specify the protocol and endpoint to receive notifications
	protocol := "email"
	endpoint := "abc@gmail.com"

	// Subscribe to the SNS topic
	subscriptionInput := &sns.SubscribeInput{
		TopicArn: aws.String(topicARN),
		Protocol: aws.String(protocol),
		Endpoint: aws.String(endpoint),
	}

	subscriptionOutput, err := svc.Subscribe(subscriptionInput)
	if err != nil {
		fmt.Println("Failed to subscribe to SNS topic:", err)
		return
	}

	fmt.Println("Successfully subscribed to SNS topic!")
	fmt.Println("Subscription ARN:", *subscriptionOutput.SubscriptionArn)
}
