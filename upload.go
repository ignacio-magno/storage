package storageS3

import (
	"bytes"
	"context"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// UploadObject in metadata, set values optionals how "Content-Type" or "Content-Disposition"
func UploadObject(u IStorageS3, data multipart.File, metadata func(*s3.PutObjectInput)) error {

	var putObjectInput s3.PutObjectInput

	metadata(&putObjectInput)

	putObjectInput.Bucket = aws.String(u.GetBucket())
	putObjectInput.Body = data
	putObjectInput.Key = aws.String(u.GetKey())

	_, err := client.PutObject(context.Background(), &putObjectInput)

	return err
}

// UploadObject in metadata, set values optionals how "Content-Type" or "Content-Disposition"
func UploadObjectFromBytes(u IStorageS3, data []byte, metadata func(*s3.PutObjectInput)) error {

	var putObjectInput s3.PutObjectInput

	metadata(&putObjectInput)

	file := bytes.NewReader(data)

	putObjectInput.Bucket = aws.String(u.GetBucket())
	putObjectInput.Body = file
	putObjectInput.Key = aws.String(u.GetKey())

	_, err := client.PutObject(context.Background(), &putObjectInput)

	return err
}
