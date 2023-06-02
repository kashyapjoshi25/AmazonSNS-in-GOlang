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
		Region:      aws.String("ap-south-1"), // Replace with your desired AWS region
		Credentials: credentials.NewStaticCredentials("AWS ACCESS KEY", "AWS ACCESS KEY ID", ""),
	})
	if err != nil {
		fmt.Println("Failed to create AWS session:", err)
		return
	}

	// Create an SNS client
	svc := sns.New(sess)

	// Specify the topic name
	topicName := "Testing_SNS_2"

	// Create the SNS topic
	createTopicInput := &sns.CreateTopicInput{
		Name: aws.String(topicName),
	}
	createTopicOutput, err := svc.CreateTopic(createTopicInput)
	if err != nil {
		fmt.Println("Failed to create topic:", err)
		return
	}

	// Get the ARN of the created topic
	topicARN := aws.StringValue(createTopicOutput.TopicArn)
	fmt.Println("Topic ARN:", topicARN)
}
