package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type PresignedClient struct {
	client   *s3.PresignClient
	bucket   string
	key      string
	duration time.Duration
}

func GetPresignedClient() PresignedClient {
	presigned := s3.NewPresignClient(client)
	return PresignedClient{client: presigned}
}

func (p *PresignedClient) SetBucket(key string) {
	p.bucket = key
}

func (p *PresignedClient) SetKey(key string) {
	p.key = key
}

func (p *PresignedClient) SetTime(val time.Duration) {
	p.duration = val
}

// return the url
func (p *PresignedClient) GeneratePresignedUrlGet() (string, error) {

	res, err := p.client.PresignGetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(p.bucket),
		Key:    aws.String(p.key),
	}, func(o *s3.PresignOptions) {
		o.Expires = p.duration
	})

	return res.URL, err
}

// return url to put object
func (p *PresignedClient) GeneratePresignedUrlPut() {

	res, err := p.client.PresignPutObject(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(p.bucket),
		Key:    aws.String(p.key),
	}, func(o *s3.PresignOptions) {
		o.Expires = p.duration
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("res.Method: %v\n", res.Method)
	for k, v := range res.SignedHeader {
		fmt.Printf("k: %v\n", k)
		fmt.Printf("v: %v\n", v)
	}

	fmt.Printf("res.URL: %v\n", res.URL)

}
