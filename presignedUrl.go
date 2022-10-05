package storageS3

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GeneratePresignedUrlGet(p IStorageS3, options func(options *s3.PresignOptions)) (string, error) {

	res, err := presignedClient.PresignGetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(p.GetBucket()),
		Key:    aws.String(p.GetKey()),
	}, options)

	if err != nil {
		return "", err
	} else {
		return res.URL, err
	}
}

func GeneratePresignedUrlPut(p IStorageS3, options func(presignOptions *s3.PresignOptions)) (string, error) {

	res, err := presignedClient.PresignPutObject(context.Background(),
		&s3.PutObjectInput{
			Bucket: aws.String(p.GetBucket()),
			Key:    aws.String(p.GetKey()),
		}, options)

	if err != nil {
		return "", err
	} else {
		return res.URL, err
	}
}
