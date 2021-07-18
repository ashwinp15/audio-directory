package graph

import (
	"context"
	"log"

	"github.com/99designs/gqlgen/graphql"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var s3client *s3.Client

const (
	BUCKET_NAME string = "nooble-bucket"
	REGION             = "us-east-1"
)

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("coolash"), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	s3client = s3.NewFromConfig(cfg)

}

func ListBuckets() (resp *s3.ListBucketsOutput) {
	resp, err := s3client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		log.Println(err)
	}
	return resp
}

func UploadObject(obj graphql.Upload) (resp *s3.PutObjectOutput) {
	input := &s3.PutObjectInput{
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(obj.Filename),
		Body:   obj.File,
	}

	resp, err := s3client.PutObject(context.TODO(), input)
	if err != nil {
		log.Println("Couldn't upload object", err)
	}
	return resp
}
