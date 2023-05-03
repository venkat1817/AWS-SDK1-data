// Go code that retrieves information about various AWS resources across all regions: services:-ec2,rds,s3

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
	// Get a list of all regions
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	ec2svc := ec2.New(sess)
	regionOutput, err := ec2svc.DescribeRegions(&ec2.DescribeRegionsInput{})
	if err != nil {
		fmt.Println("Error getting regions: ", err)
		return
	}

	for _, region := range regionOutput.Regions {
		fmt.Printf("Region: %s\n", *region.RegionName)

		// Get information about S3 buckets
		s3svc := s3.New(sess, &aws.Config{Region: region.RegionName})
		bucketOutput, err := s3svc.ListBuckets(&s3.ListBucketsInput{})
		if err != nil {
			fmt.Println("Error getting S3 buckets: ", err)
			return
		}
		fmt.Println("S3 Buckets:")
		for _, bucket := range bucketOutput.Buckets {
			fmt.Printf("\t%s\n", *bucket.Name)
		}

		// Get information about RDS instances
		rdssvc := rds.New(sess, &aws.Config{Region: region.RegionName})
		rdsOutput, err := rdssvc.DescribeDBInstances(&rds.DescribeDBInstancesInput{})
		if err != nil {
			fmt.Println("Error getting RDS instances: ", err)
			return
		}
		fmt.Println("RDS Instances:")
		for _, dbInstance := range rdsOutput.DBInstances {
			fmt.Printf("\t%s\n", *dbInstance.DBInstanceIdentifier)
		}

		// Get information about EC2 instances
		ec2svc := ec2.New(sess, &aws.Config{Region: region.RegionName})
		ec2Output, err := ec2svc.DescribeInstances(&ec2.DescribeInstancesInput{})
		if err != nil {
			fmt.Println("Error getting EC2 instances: ", err)
			return
		}
		fmt.Println("EC2 Instances:")
		for _, reservation := range ec2Output.Reservations {
			for _, instance := range reservation.Instances {
				fmt.Printf("\t%s\n", *instance.InstanceId)
			}
		}
		fmt.Println()
	}
}
