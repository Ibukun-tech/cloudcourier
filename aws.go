package cloudcourier

import (
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type AwsManufacture struct {
	Region string
	Bucket string
}

func (a *AwsManufacture) GetProvider() cloudCourierProvider {
	return AWS
}
func init() {
	storeFunc[AWS] = newAWSClient
}

type AWSClient struct {
	AwConfig   *AwsManufacture
	BucketName string
	S3         *s3.S3
}

func newAWSClient(cbb Provider) (StorageClient, error) {
	awConfig, ok := cbb.(*AwsManufacture)
	if !ok {
		return nil, fmt.Errorf("incorrect configuration")
	}
	sess, err := session.NewSession(&aws.Config{
		Region: &awConfig.Region,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS session: %v", err)
	}
	s3Svc := s3.New(sess)
	return &AWSClient{S3: s3Svc,
		BucketName: awConfig.Bucket, AwConfig: awConfig}, nil
}

// AWSClient struct encapsulates the AWS S3 client to manage storage operations.
// type AWSClient struct {
// 	S3         *s3.S3 // The S3 service client
// 	BucketName string // The name of the bucket to operate on
// }

// NewAWSClient creates a new client for AWS S3 operations.

// It requires AWS credentials and a region to initialize the S3 service client.

// UploadFile uploads a file to S3 from an io.Reader.
func (client *AWSClient) UploadFile(filePath string, reader io.Reader) error {
	uploader := s3manager.NewUploaderWithClient(client.S3)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(client.BucketName),
		Key:    aws.String(filePath),
		Body:   reader,
	})
	if err != nil {
		return fmt.Errorf("failed to upload file to S3: %v", err)
	}
	return nil
}

// DeleteFile deletes a file from S3 by its key.
func (client *AWSClient) DeleteFile(fileID string) error {
	_, err := client.S3.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(client.BucketName),
		Key:    aws.String(fileID),
	})
	if err != nil {
		return fmt.Errorf("failed to delete file from S3: %v", err)
	}
	return nil
}

// ListFiles lists all files in a specified directory of the S3 bucket.
func (client *AWSClient) ListFiles(directory string) ([]string, error) {
	var fileNames []string
	err := client.S3.ListObjectsV2Pages(&s3.ListObjectsV2Input{
		Bucket: aws.String(client.BucketName),
		Prefix: aws.String(directory),
	}, func(page *s3.ListObjectsV2Output, lastPage bool) bool {
		for _, obj := range page.Contents {
			fileNames = append(fileNames, *obj.Key)
		}
		return true // continue paging
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list files in S3: %v", err)
	}
	return fileNames, nil
}

// GetFile retrieves a file as an io.Reader by its key from S3.
func (client *AWSClient) GetFile(fileID string) (io.Reader, error) {
	output, err := client.S3.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(client.BucketName),
		Key:    aws.String(fileID),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve file from S3: %v", err)
	}
	return output.Body, nil
}
