package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var PGclient *pgxpool.Pool

func init() {
	conn, err := pgxpool.Connect(context.TODO(), os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal(err)
	}
	PGclient = conn
	fmt.Println("Connected to DB")
}
