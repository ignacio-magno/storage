package storage

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// =====================================================================================================================
type Upload struct {
	base
	data []byte
}

func (j Upload) Upload(bucket string, key string, data []byte) *Upload {
	var up Upload
	up.constructorBase(bucket, key)
	up.data = data
	return &up
}

func (s *Upload) getData() []byte {
	return s.data
}

func (j *Upload) UploadObject() error {
	return uploadObject(j)
}

// without struct

func uploadObject(u IUpload) error {

	reader := bytes.NewBuffer(u.getData())

	_, err := client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(u.getBucket()),
		Key:    aws.String(u.getKey()),
		Body:   reader,
	})

	return err
}
