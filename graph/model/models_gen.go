// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/99designs/gqlgen/graphql"
)

type Creator struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type NewCreator struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewNooble struct {
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Category    string         `json:"category"`
	File        graphql.Upload `json:"file"`
	Creator     string         `json:"creator"`
}

type Nooble struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Category    string   `json:"category"`
	Audio       string   `json:"audio"`
	Creator     *Creator `json:"creator"`
}
