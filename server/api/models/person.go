package models

import "github.com/graphql-go/graphql"

type Person struct {
	Job  string   `json:"job"`
	Role []string `json:"role"`
	Name string   `json:"name"`
}

var PersonType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Person",
	Fields: graphql.Fields{
		"Job": &graphql.Field{
			Type: graphql.String,
		},
		"Role": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
		"Name": &graphql.Field{
			Type: graphql.String,
		},
	},
})
