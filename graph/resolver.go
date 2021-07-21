package graph

import (
	"context"
	"fmt"
	"log"

	"github.com/99designs/gqlgen/graphql"
	"github.com/ashwinp15/audio-directory/database"
	"github.com/ashwinp15/audio-directory/graph/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	//"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

var s3client *s3.Client
var PGclient *pgxpool.Pool

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

	PGclient = database.PGclient

}

type Resolver struct {
	nooble     *model.Nooble
	noobleList []*model.Nooble
}

func (r Resolver) ReadAllNoobles() ([]*model.Nooble, error) {
	query := fmt.Sprintf(`
SELECT
	title,
	description,
	category,
	audio
	FROM public.noobles`)
	rows, err := PGclient.Query(context.TODO(), query)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var nooble model.Nooble
	for rows.Next() {
		if err := rows.Scan(&nooble.Title, &nooble.Description, &nooble.Category, &nooble.Audio); err != nil {
			log.Println(err)
			return nil, err
		}
		fmt.Printf("received entry: %v\n", nooble)
		r.noobleList = append(r.noobleList, &nooble)
	}
	return r.noobleList, nil
}

func (r Resolver) ReadSingleNooble(id string) (*model.Nooble, error) {
	sql := fmt.Sprintf(`
SELECT n.id, n.title, n.category, n.description, n.audio, c.email, c.name
	 FROM noobles n INNER JOIN creators c
	 ON n.creator = c.email
	 WHERE n.id = $1
	 `)
	var nooble model.Nooble
	var creator model.Creator
	row := database.PGclient.QueryRow(context.TODO(), sql, id)
	fmt.Println("row read successfully")
	if err := row.Scan(&nooble.ID, &nooble.Title, &nooble.Category,
		&nooble.Description, &nooble.Audio, &creator.Email, &creator.Name); err != nil {
		return nil, err
	}
	nooble.Creator = &creator
	fmt.Println("row scanned successfully")
	return &nooble, nil
}

func (r Resolver) PutNooble(obj graphql.Upload) {
	r.UploadAudio(obj)
	r.addToDB()
}

func (r Resolver) addToDB() {
	query := fmt.Sprintf(
		`INSERT INTO noobles (title, category, description, audio, creator)
VALUES ($1, $2, $3, $4, $5)`,
	)

	commandTag, err := PGclient.Exec(context.TODO(), query,
		r.nooble.Title, r.nooble.Category, r.nooble.Description, r.nooble.Audio, r.nooble.Creator.Email)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("command Tag ", commandTag)
}

func (r Resolver) UploadAudio(obj graphql.Upload) {
	input := &s3.PutObjectInput{
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(obj.Filename),
		Body:   obj.File,
	}

	resp, err := s3client.PutObject(context.TODO(), input)
	if err != nil {
		log.Println("Couldn't upload object", err)
	}
	fmt.Println("upload response: ", resp)
}
