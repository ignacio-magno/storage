package storage

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// =====================================================================================================================
type MoveObject struct {
	base
	sourceKey string
}

func (m MoveObject) MoveObject(bucket, key, sourceKey string) *MoveObject {
	var mo MoveObject
	mo.constructorBase(bucket, key)
	mo.sourceKey = sourceKey
	return &mo
}

func (m *MoveObject) CopyObject() error {
	_, err := client.CopyObject(context.Background(), &s3.CopyObjectInput{
		Bucket:     aws.String(m.getBucket()),
		Key:        aws.String(m.getKey()),
		CopySource: aws.String(m.sourceKey),
	})

	return err
}
