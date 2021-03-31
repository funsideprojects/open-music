package main

import (
	"fsp/open-music/identity/graph"
	"fsp/open-music/identity/graph/generated"

	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const defaultPort = "8080"

func main() {
	e := echo.New()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	e.POST("/gql", func(context echo.Context) error {
		srv.ServeHTTP(context.Response(), context.Request())
		return nil
	})

	e.GET("/playground", func(c echo.Context) error {
		playground.Handler("GraphQL", "/query").ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Welcome!")
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	e.Logger.Fatal(e.Start(":" + port))
}
