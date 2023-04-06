package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jaimy-monsuur/movie-api/src/config"
	_ "github.com/jaimy-monsuur/movie-api/src/docs"
	"github.com/jaimy-monsuur/movie-api/src/middlewares"
	"github.com/jaimy-monsuur/movie-api/src/routes"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
}

func main() {

	router := gin.Default()

	router.Use(middlewares.CorsMiddleware())

	routes.IndexRoutes(router)

	routes.UserRoutes(router)

	routes.AuthRoutes(router)

	routes.MovieRoutes(router)

	routes.ReviewRoutes(router)

	router.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run()
}
