package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/jackc/pgx/v4/pgxpool"
)

var PGclient *pgxpool.Pool
var S3client *s3.Client

func init() {
	conn, err := pgxpool.Connect(context.TODO(), os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal(err)
	}
	PGclient = conn
	fmt.Println("Connected to DB")

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("coolash"), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	S3client = s3.NewFromConfig(cfg)
}
