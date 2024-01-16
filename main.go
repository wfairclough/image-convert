package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

var (
	invokeCount int
	myObjects   []types.Object
)

func init() {

	log = logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
  
  log.Info("Initializing Lambda")

	// Load the SDK configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ca-central-1"))
	if err != nil {
		log.Fatalf("Unable to load SDK config: %v", err)
	}

	// Initialize an S3 client
	svc := s3.NewFromConfig(cfg)

	// Define the bucket name as a variable so we can take its address
	bucketName := "willf-dev"
	input := &s3.ListObjectsV2Input{
		Bucket: &bucketName,
	}

	// List objects in the bucket
	result, err := svc.ListObjectsV2(context.TODO(), input)
	if err != nil {
		log.Fatalf("Failed to list objects: %v", err)
	}
	myObjects = result.Contents
  log.Info("Objects listed")
}

func LambdaHandler(ctx context.Context, event events.S3Event) (int, error) {
  log.Info("Lambda invoked")

  // log the event as json 
  log.WithFields(logrus.Fields{
    "event": event,
  }).Info("Lambda invoked")

	invokeCount++
	for i, obj := range myObjects {
		log.Printf("object[%d] size: %d key: %s", i, obj.Size, *obj.Key)
	}
  log.Info("Lambda completed")
	return invokeCount, nil
}

func main() {
	lambda.Start(LambdaHandler)
}
