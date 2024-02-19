package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AWSClient struct {
	S3Client *s3.Client
}

var awsClient *AWSClient

func initAWS() error {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		return fmt.Errorf("couldn't load default configuration. Have you set up your AWS account?")
	}
	s3Client := s3.NewFromConfig(sdkConfig)
	awsClient = &AWSClient{
		S3Client: s3Client,
	}
	return nil
}

func (awsClient AWSClient) UploadFile(bucketName string, objectKey string, fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		printf("Couldn't open file %v to upload. Here's why: %v\n", fileName, err)
	} else {
		defer file.Close()
		contentType, err := GetFileContentType(file)
		if err != nil {
			println("Can't find context type")
		}
		_, err = awsClient.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket:      aws.String(bucketName),
			Key:         aws.String(objectKey),
			Body:        file,
			ContentType: aws.String(contentType),
		})
		if err != nil {
			printf("Couldn't upload file %v to %v:%v. Here's why: %v\n",
				fileName, bucketName, objectKey, err)
		}
	}
	return err
}
