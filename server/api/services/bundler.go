package services

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"getAllMovies":               getAllMovies,
		"movieByTitle":               getMovieByTitle,
		"moviesWithinThreeRelations": moviesWithinThreeRelations,
		"moviesByDirector":           moviesByDirector,
		"moviesByActor":              moviesByActor,
	},
})

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
})

func RunServer() {
	router := gin.Default()
	setupCors(router)

	handler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	router.GET("/graphql", gin.WrapH(handler))
	router.POST("/graphql", gin.WrapH(handler))

	router.Run(":8080")
}

func setupCors(router gin.IRouter) {
	router.Use(cors.New(cors.Config{
		AllowHeaders:     []string{"content-type"},
		AllowCredentials: true,
		AllowMethods:     []string{"POST", "GET"},
		AllowOrigins:     []string{"http://127.0.0.1"},
	}))
}
