package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
	"github.com/go-chi/chi/middleware"
	"graphQL/gql"
	"graphQL/postgres"
	"graphQL/server"
	"log"
	"net/http"
)

func main()  {
	router, db := Iniatilize()
	defer db.Close()

	log.Fatal(http.ListenAndServe(":4000", router))
}

func Iniatilize() (*chi.Mux, *postgres.DB)  {
	router := chi.NewRouter()

	db, err := postgres.New(postgres.ConnString("localhost", 5432, "postgres", "postgres", "coba"),)
	if err != nil{
		fmt.Println("masuk")
		log.Fatal(err)
	}

	rootQuery := gql.NewRoot(db)
	sc, err := graphql.NewSchema(graphql.SchemaConfig{Query:rootQuery.Query})
	if err != nil{
		fmt.Println("error creating schema: ",err)
	}

	s := server.Server{
		GqlSchema:&sc,
	}
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.StripSlashes,
		middleware.Recoverer,)

	router.Get("/graphql", s.GraphQl())

	return router, db
}