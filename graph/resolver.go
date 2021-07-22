package graph

import (
	"context"
	"fmt"
	"log"

	"github.com/99designs/gqlgen/graphql"
	"github.com/ashwinp15/audio-directory/database"
	"github.com/ashwinp15/audio-directory/graph/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

const (
	BUCKET_NAME string = "nooble-bucket"
	REGION             = "us-east-1"
)

type Resolver struct {
	nooble     *model.Nooble
	noobleList []*model.Nooble
}

func (r Resolver) ReadAllNoobles() ([]*model.Nooble, error) {
	query := fmt.Sprintf(`
SELECT
	 id,
	title,
	description,
	category,
	audio
	FROM public.noobles`)
	rows, err := database.PGclient.Query(context.TODO(), query)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var nooble model.Nooble
	for rows.Next() {
		if err := rows.Scan(&nooble.ID, &nooble.Title, &nooble.Description, &nooble.Category, &nooble.Audio); err != nil {
			log.Println(err)
			return nil, err
		}
		fmt.Printf("received entry: %v\n", nooble)
		r.noobleList = append(r.noobleList, &nooble)
	}
	return r.noobleList, nil
}

func (r Resolver) ReadSingleNooble(id string) (*model.Nooble, error) {
	sql := fmt.Sprintf(
		`
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

// Updating nooble
func (r Resolver) UpdateDetails(input *model.UpdateNooble) (*string, error) {
	sql := fmt.Sprintf(
		`
	 UPDATE noobles SET
	 title = COALESCE($1, title),
	 category = COALESCE($2, category),
	 description = COALESCE($3, description)
	 WHERE id = $4
	 `)

	commandTag, err := database.PGclient.Exec(context.TODO(), sql,
		input.Title, input.Category, input.Description, r.nooble.ID)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%v\n", commandTag)
	return &r.nooble.ID, nil
}

func (r Resolver) Delete() (*string, error) {
	readQuery := fmt.Sprintf(
		`SELECT audio FROM noobles
			WHERE id = $1`,
	)
	var filename string
	if err := database.PGclient.QueryRow(context.TODO(), readQuery, r.nooble.ID).Scan(&filename); err != nil {
		log.Println("Postgres read error: ", err)
		return nil, err
	}
	if _, err := database.S3client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(filename),
	}); err != nil {
		log.Println("S3 delete object error: ", err)
		return nil, err
	}

	deleteQuery := fmt.Sprintf(
		`DELETE FROM noobles
			WHERE id = $1`,
	)
	if _, err := database.PGclient.Exec(context.TODO(), deleteQuery, r.nooble.ID); err != nil {
		log.Println("Postgres delete error", err)
		return nil, err
	}
	return &r.nooble.ID, nil
}

// Creating new nooble
func (r Resolver) PutNooble(obj graphql.Upload) {
	r.UploadAudio(obj)
	r.addToDB()
}

// Helper functions for PutNooble
func (r Resolver) addToDB() {
	query := fmt.Sprintf(
		`INSERT INTO noobles (title, category, description, audio, creator)
VALUES ($1, $2, $3, $4, $5)`,
	)

	commandTag, err := database.PGclient.Exec(context.TODO(), query,
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

	resp, err := database.S3client.PutObject(context.TODO(), input)
	if err != nil {
		log.Println("Couldn't upload object", err)
	}
	fmt.Println("upload response: ", resp)
}
