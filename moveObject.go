package storageS3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type IMoveObject interface {
	IStorageS3
	GetSourceKey() string
}

func CopyObject(m IMoveObject) error {
	_, err := client.CopyObject(context.Background(), &s3.CopyObjectInput{
		Bucket:     aws.String(m.GetBucket()),
		Key:        aws.String(m.GetKey()),
		CopySource: aws.String(m.GetSourceKey()),
	})

	return err
}
