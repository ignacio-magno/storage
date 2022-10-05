package storageS3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func DeleteObject(d IStorageS3) error {
	_, err := client.DeleteObject(context.Background(), &s3.DeleteObjectInput{
		Bucket: aws.String(d.GetBucket()),
		Key:    aws.String(d.GetKey()),
	})

	return err
}
