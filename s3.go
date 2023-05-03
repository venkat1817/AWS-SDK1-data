// this code is for to get data from  only in one region services:-s3,rds,ec2
package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	// Create a new session in the us-east-1 region.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	if err != nil {
		fmt.Println("Error creating session: ", err)
		return
	}

	// Get EC2 instances.
	ec2Svc := ec2.New(sess)
	ec2Result, err := ec2Svc.DescribeInstances(nil)
	if err != nil {
		fmt.Println("Error describing EC2 instances: ", err)
		return
	}

	// Get RDS instances.
	rdsSvc := rds.New(sess)
	rdsResult, err := rdsSvc.DescribeDBInstances(nil)
	if err != nil {
		fmt.Println("Error describing RDS instances: ", err)
		return
	}

	// Get S3 buckets.
	s3Svc := s3.New(sess)
	s3Result, err := s3Svc.ListBuckets(nil)
	if err != nil {
		fmt.Println("Error listing S3 buckets: ", err)
		return
	}

	// Print EC2 instances.
	fmt.Println("EC2 instances:")
	for _, reservation := range ec2Result.Reservations {
		for _, instance := range reservation.Instances {
			fmt.Printf("ID: %s, State: %s\n", *instance.InstanceId, *instance.State.Name)
		}
	}

	// Print RDS instances.
	fmt.Println("RDS instances:")
	for _, instance := range rdsResult.DBInstances {
		fmt.Printf("ID: %s, Engine: %s, Status: %s\n", *instance.DBInstanceIdentifier, *instance.Engine, *instance.DBInstanceStatus)
	}

	// Print S3 buckets.
	fmt.Println("S3 buckets:")
	for _, bucket := range s3Result.Buckets {
		locationResult, err := s3Svc.GetBucketLocation(&s3.GetBucketLocationInput{
			Bucket: bucket.Name,
		})
		if err != nil {
			fmt.Printf("Error getting location for bucket %s: %s\n", *bucket.Name, err)
		}
		fmt.Printf("Name: %s, CreationDate: %s, Location: %s\n", *bucket.Name, bucket.CreationDate, *locationResult.LocationConstraint)
	}
}
