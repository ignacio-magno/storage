package storageS3

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// DownloadObject without struct
func DownloadObject(d IStorageS3) ([]byte, error) {
	res, err := client.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(d.GetBucket()),
		Key:    aws.String(d.GetKey()),
	})

	if err != nil {
		return nil, err
	}

	return io.ReadAll(res.Body)
}
