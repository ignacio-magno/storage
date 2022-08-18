package storage

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// ============================================================================================================================

type Download struct {
	base
}

func (d Download) Download(bucket string, key string) *Download {
	var dow Download
	dow.constructorBase(bucket, key)
	return &dow
}

func (d *Download) DownloadObject() ([]byte, error) {
	return downloadObject(d)
}

//  without struct
func downloadObject(d IDownload) ([]byte, error) {
	res, err := client.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(d.getBucket()),
		Key:    aws.String(d.getKey()),
	})

	if err != nil {
		return nil, err
	}

	return io.ReadAll(res.Body)
}
