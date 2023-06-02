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
		Region:      aws.String("ap-south-1"),                                                                                 // Replace with your desired AWS region
		Credentials: credentials.NewStaticCredentials("AKIAVGIMFRGDZUUMGGHW", "43GGJBJFsfdh342hPgDXf1IcG5V1+s3NrW+v2zDw", ""), // Uncomment if not using shared credentials file
	})
	if err != nil {
		fmt.Println("Failed to create session:", err)
		return
	}

	// Create an SNS client
	svc := sns.New(sess)

	// Specify the topic ARN you want to send the message to
	topicARN := "arn:aws:sns:ap-south-1:357044685191:Testing_SNS"

	// Specify the message you want to send
	message := "Hello, subscribers from kashyap!"

	// Send the message to the SNS topic
	publishInput := &sns.PublishInput{
		Message:  aws.String(message),
		TopicArn: aws.String(topicARN),
	}

	publishOutput, err := svc.Publish(publishInput)
	if err != nil {
		fmt.Println("Failed to send message to SNS topic:", err)
		return
	}

	fmt.Println("Message sent to SNS topic!")
	fmt.Println("Message ID:", *publishOutput.MessageId)
}
