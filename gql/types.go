package gql

import "github.com/graphql-go/graphql"

var Program = graphql.NewObject(
	graphql.ObjectConfig{
		Name:"Program",
		Fields: graphql.Fields{
			"id": &graphql.Field{Type:graphql.Int,},
			"created": &graphql.Field{Type:graphql.DateTime},
			"createdby": &graphql.Field{Type:graphql.String},
			"modified": &graphql.Field{Type:graphql.DateTime},
			"modifiedby":&graphql.Field{Type:graphql.String},
			"active":&graphql.Field{Type:graphql.Boolean},
			"isdeleted":&graphql.Field{Type:graphql.Boolean},
			"deleted":&graphql.Field{Type:graphql.DateTime},
			"deletedby":&graphql.Field{Type:graphql.String},
			"programname":&graphql.Field{Type:graphql.String},
			"programimage":&graphql.Field{Type:graphql.String},
			"programstartdate":&graphql.Field{Type:graphql.DateTime},
			"programenddate":&graphql.Field{Type:graphql.DateTime},
			"programdescription":&graphql.Field{Type:graphql.String},
			"card":&graphql.Field{Type:graphql.String},
			"ouletid":&graphql.Field{Type:graphql.Int},
			"merchantname":&graphql.Field{Type:graphql.String},
			"categoryid":&graphql.Field{Type:graphql.Int},
			"benefit":&graphql.Field{Type:graphql.String},
			"termsandcondition":&graphql.Field{Type:graphql.String},
			"tier":&graphql.Field{Type:graphql.String},
			"redemrules":&graphql.Field{Type:graphql.String},
			"rewardtarget":&graphql.Field{Type:graphql.Float},
			"qrcodeid":&graphql.Field{Type:graphql.String},
		},
	})
