package models

import "github.com/graphql-go/graphql"

type Movie struct {
	Released int64    `json:"released"`
	Title    string   `json:"title,omitempty"`
	Tagline  string   `json:"tagline,omitempty"`
	Votes    int64    `json:"votes,omitempty"`
	Cast     []Person `json:"cast,omitempty"`
}

type MovieResult struct {
	Movie `json:"movie"`
}

var MovieType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Movie",
	Fields: graphql.Fields{
		"Released": &graphql.Field{
			Type: graphql.Int,
		},
		"Title": &graphql.Field{
			Type: graphql.String,
		},
		"Tagline": &graphql.Field{
			Type: graphql.String,
		},
		"Votes": &graphql.Field{
			Type: graphql.Int,
		},
		"Cast": &graphql.Field{
			Type: PersonType,
		},
	},
})
