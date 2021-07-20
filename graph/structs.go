package graph

import (
	"github.com/99designs/gqlgen/graphql"
)

type Creator struct {
	ID    string `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`
}

type NewCreator struct {
	ID    string `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`
}

type NewNooble struct {
	Title       string         `json:"title" db:"title"`
	Description string         `json:"description" db:"description"`
	Category    string         `json:"category" db:"category"`
	File        graphql.Upload `json:"file" db:"file"`
}

type Nooble struct {
	ID          string   `json:"id" db:"id"`
	Title       string   `json:"title" db:"title"`
	Description string   `json:"description" db:"description"`
	Category    string   `json:"category" db:"category"`
	Audio       string   `json:"audio" db:"audio"`
	Creator     *Creator `json:"creator" db:"creator"`
}
