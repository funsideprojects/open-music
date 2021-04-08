package main

import (
	"log"

	"fsp/open-music/identity/graph"
	"fsp/open-music/identity/graph/generated"

	"fsp/open-music/packages/colors"
	"fsp/open-music/packages/database"
	"fsp/open-music/packages/env"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load .env file
	env.LoadEnvConfigs()
	port := env.GetEnv("IDENTITY_PORT")

	// Connect to database
	database.MongoConnection()

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	// ? Middlewares
	// Limit body size - https://echo.labstack.com/middleware/body-limit
	e.Use(middleware.BodyLimit("2MB"))

	// Cors - https://echo.labstack.com/middleware/cors
	corsWhitelist := env.GetEnv("WEB_URL")
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{corsWhitelist},
	}))

	// CSRF - https://echo.labstack.com/middleware/csrf

	// Request Rate limit - https://echo.labstack.com/middleware/rate-limiter
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	// Tags to construct the logger format - https://echo.labstack.com/middleware/logger
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method = ${method}, status = ${status}, latency = ${latency_human}\n",
	}))

	// Recover from panic - https://echo.labstack.com/middleware/recover
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         1 << 10, // 1 KB
		DisableStackAll:   false,
		DisablePrintStack: false,
	}))

	// ? GraphQL
	graphqlServer := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	e.POST("/identity", func(context echo.Context) error {
		graphqlServer.ServeHTTP(context.Response(), context.Request())
		return nil
	})

	e.GET("/pg-identity", func(c echo.Context) error {
		playground.Handler("GraphQL", "/identity").ServeHTTP(c.Response(), c.Request())
		return nil
	})

	// e.GET("/", func(context echo.Context) error {
	// 	return context.String(http.StatusOK, "Welcome!")
	// })

	log.Println(colors.Success("[Identity]")+" Server started at port:", port)
	e.Logger.Fatal(e.Start(":" + port))
}
