package storage

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Delete struct {
	bucket string
	key    string
}

func (d Delete) Delete(bucket, key string) *Delete {
	var del Delete
	del.bucket = bucket
	del.key = key
	return &del
}

func (d *Delete) DeleteObject() error {
	_, err := client.DeleteObject(context.Background(), &s3.DeleteObjectInput{
		Bucket: aws.String(d.bucket),
		Key:    aws.String(d.key),
	})

	return err
}
