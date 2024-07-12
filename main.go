package main

import (
	"context"

	s3api "github.com/thoratvinod/Upwork-BookingAppPortfolio/s3"
	twilioapi "github.com/thoratvinod/Upwork-BookingAppPortfolio/twilio"
)

func main() {

	ctx := context.Background()
	s3api := s3api.S3API{
		Region:     "us-east-1",
		BucketName: "book-seling-app",
	}

	// upload file
	s3api.UploadFileToS3(ctx, "demo", "demo.txt")

	// download file
	s3api.UploadFileToS3(ctx, "demo", "demo.txt")

	req := twilioapi.SMSRequest{
		From: "sender-number",
		To:   "receivers-number",
		Body: "Sample body",
	}
	// send sms
	twilioapi.SendSMS(&req)
}
