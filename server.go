package main

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pwnd27/go_app/app"
	"github.com/pwnd27/go_app/db"
	"github.com/pwnd27/go_app/graph"
	"os"
)

func main() {
	// Setting up Gin
	r := gin.Default()
	r.Use(app.GinContextToContextMiddleware())

	dbPool, err := pgxpool.New(context.Background(), "postgres://user:pass@localhost:5432/mydb?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbPool.Close()

	store := db.NewStore(dbPool)
	r.POST("/query", graphqlHandler(store))
	r.GET("/", playgroundHandler())
	r.Run()

}

// Defining the Graphql handler
func graphqlHandler(store db.Store) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DB: store}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
