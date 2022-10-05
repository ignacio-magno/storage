package storageS3

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var client *s3.Client
var presignedClient *s3.PresignClient

func init() {

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		panic(err)
		return
	}

	client = s3.NewFromConfig(cfg)

	presignedClient = s3.NewPresignClient(client)
}
