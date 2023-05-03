//this is code is only to get s3 bucket data from all regions services:-s3

package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
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

	// Create an S3 service client.
	svc := s3.New(sess)

	// Get a list of all S3 buckets in the account.
	result, err := svc.ListBuckets(nil)
	if err != nil {
		fmt.Println("Error listing buckets: ", err)
		return
	}

	// Print details of each S3 bucket.
	fmt.Println("S3 Buckets:")
	for _, b := range result.Buckets {
		// Get the location of the S3 bucket.
		locInput := &s3.GetBucketLocationInput{
			Bucket: aws.String(*b.Name),
		}
		location, err := svc.GetBucketLocation(locInput)
		if err != nil {
			fmt.Println("Error getting bucket location: ", err)
			return
		}

		fmt.Printf("Name: %s, CreationDate: %s, Location: %s\n", *b.Name, *b.CreationDate, *location.LocationConstraint)
	}
}
