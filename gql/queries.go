package gql

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"graphQL/postgres"
)

type Root struct {
	Query 	*graphql.Object
}

func NewRoot(db *postgres.DB) *Root  {
	resolver := Resolver{db:db}

	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"program": &graphql.Field{
						Type: graphql.NewList(Program),
						Args: graphql.FieldConfigArgument{
							"page": &graphql.ArgumentConfig{
								Type: graphql.Int,
								DefaultValue:0,
							},
							"size": &graphql.ArgumentConfig{
								Type: graphql.Int,
								DefaultValue:0,
							},
							"sort": &graphql.ArgumentConfig{
								Type: graphql.Int,
								DefaultValue:0,
							},
							"category": &graphql.ArgumentConfig{
								Type: graphql.Int,
								DefaultValue:0,
							},
							"id": &graphql.ArgumentConfig{
								Type: graphql.Int,
								DefaultValue:0,
							},

						},
						Resolve: resolver.SpecialProgramResolver,
					},
				},
			}),
	}
	fmt.Println(root)
	return &root
}
