package server

import (
	"encoding/json"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
	"graphQL/gql"

	"net/http"
)

type Server struct {
	GqlSchema 	*graphql.Schema
}

type reqBody struct {
	Query string `json:"query"`
}

func (s *Server) GraphQl() http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil{
			http.Error(w, "Must provide graph query in request body", 400)
			return
		}
		var rBody reqBody

		err := json.NewDecoder(r.Body).Decode(&rBody)
		if err != nil{
			http.Error(w, "Error parsing JSON request body", 400)
		}
		result := gql.ExecuteQuery(rBody.Query, *s.GqlSchema)

		render.JSON(w, r, result)
	}
}