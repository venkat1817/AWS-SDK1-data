//Go code to retrieve CloudTrail data service:-cloudtrail

package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
)

func main() {
	// Create a new session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1")},
	)
	if err != nil {
		fmt.Println("Error creating session: ", err)
		return
	}

	// Create a CloudTrail client
	svc := cloudtrail.New(sess)

	// Get the events for the past hour
	startTime := time.Now().Add(-time.Hour)
	endTime := time.Now()

	input := &cloudtrail.LookupEventsInput{
		StartTime: aws.Time(startTime),
		EndTime:   aws.Time(endTime),
	}

	result, err := svc.LookupEvents(input)
	if err != nil {
		fmt.Println("Error getting CloudTrail events: ", err)
		return
	}

	// Print the event names and event times
	fmt.Println("CloudTrail Events:")
	for _, event := range result.Events {
		fmt.Printf("%s (%s)\n", *event.EventName, event.EventTime.Format(time.RFC3339))
	}
}
