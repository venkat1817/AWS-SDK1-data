//this code retrives the autoscailing groups from the aws service:-autoscailing

package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
)

func main() {
	// it creates new session in the us-west-1 region
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	if err != nil {
		fmt.Println("Error creating session: ", err)
		return
	}

	svc := autoscaling.New(sess)

	// this is api it connects autoscailing group and it well get the data
	result, err := svc.DescribeAutoScalingGroups(&autoscaling.DescribeAutoScalingGroupsInput{})

	if err != nil {
		fmt.Printf("Error describing Auto Scaling groups: %s\n", err)
		return
	}

	// it prints all autoscailing groups and it prints minsize, maxsize of the data also
	fmt.Println("Auto Scaling Groups:")
	if len(result.AutoScalingGroups) == 0 {
		fmt.Println("No groups found.")
		return
	}
	for _, group := range result.AutoScalingGroups {
		fmt.Printf("Name: %s, MinSize: %d, MaxSize: %d\n", *group.AutoScalingGroupName, *group.MinSize, *group.MaxSize)
	}
}
