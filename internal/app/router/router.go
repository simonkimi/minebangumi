package router

import (
	graphHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/simonkimi/minebangumi/internal/app/graph"
	"github.com/simonkimi/minebangumi/internal/app/handler"
	"github.com/simonkimi/minebangumi/internal/pkg/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(middleware.ResponseWrapperMiddleware())
	r.Use(middleware.LogrusMiddleware())

	graphSrv := graphHandler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &handler.Resolver{},
	}))

	v1 := r.Group("/v1")
	graphqlGroup := v1.Group("/graphql")
	{
		graphqlGroup.POST("/query", func(c *gin.Context) {
			graphSrv.ServeHTTP(c.Writer, c.Request)
		})
		graphqlGroup.GET("/", func(c *gin.Context) {
			playground.Handler("GraphQL playground", "/graphql/query").ServeHTTP(c.Writer, c.Request)
		})
	}
	sourceGroup := v1.Group("/source")
	{
		sourceGroup.POST("/parse", handler.ParseSource)
	}
	return r
}
