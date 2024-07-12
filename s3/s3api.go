package s3api

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type S3API struct {
	Region     string
	BucketName string
}

func (s3api *S3API) UploadFileToS3(ctx context.Context, keyName, filePath string) error {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(s3api.Region))
	if err != nil {
		return fmt.Errorf("unable to load SDK config: %w", err)
	}

	s3Client := s3.NewFromConfig(cfg)

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("unable to open file: %w", err)
	}
	defer file.Close()

	_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s3api.BucketName),
		Key:    aws.String(keyName),
		Body:   file,
		ACL:    types.ObjectCannedACLPublicRead,
	})
	if err != nil {
		return fmt.Errorf("unable to upload file to S3: %w", err)
	}

	return nil
}

func (s3api *S3API) DownloadFileFromS3(keyName, downloadPath string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(s3api.Region))
	if err != nil {
		return fmt.Errorf("unable to load SDK config: %w", err)
	}

	s3Client := s3.NewFromConfig(cfg)

	getObjectInput := &s3.GetObjectInput{
		Bucket: aws.String(s3api.BucketName),
		Key:    aws.String(keyName),
	}

	resp, err := s3Client.GetObject(context.TODO(), getObjectInput)
	if err != nil {
		return fmt.Errorf("unable to download file from S3: %w", err)
	}
	defer resp.Body.Close()

	outFile, err := os.Create(downloadPath)
	if err != nil {
		return fmt.Errorf("unable to create file: %w", err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return fmt.Errorf("unable to save file: %w", err)
	}

	return nil
}
