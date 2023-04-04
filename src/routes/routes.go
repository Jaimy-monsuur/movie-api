package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jaimy-monsuur/movie-api/src/controllers"
	"github.com/jaimy-monsuur/movie-api/src/middlewares"
)

func IndexRoutes(indexRouter *gin.Engine) {
	indexRouter.GET("/", controllers.Index)
}

func UserRoutes(router *gin.Engine) {

	userRouter := router.Group("/users")

	{
		userRouter.POST("/", controllers.CreateUser)
		userRouter.GET("/", middlewares.Auth(), controllers.GetAllUsers)
		userRouter.GET("/:id", middlewares.Auth(), controllers.GetUserByID)
		userRouter.PUT("/:id", middlewares.Auth(), controllers.UpdateUser)
		userRouter.DELETE("/:id", middlewares.AdminAuth(), controllers.DeleteUser)
	}
}

func AuthRoutes(router *gin.Engine) {

	authRouter := router.Group("/auth")

	{
		authRouter.POST("/login", controllers.LoginUser)
	}
}

func MovieRoutes(router *gin.Engine) {

	movieRouter := router.Group("/movies")

	{
		movieRouter.POST("/", middlewares.AdminAuth(), controllers.CreateMovie)
		movieRouter.GET("/", middlewares.Auth(), controllers.GetAllMovies)
		movieRouter.GET("/:id", middlewares.Auth(), controllers.GetMovieByID)
		movieRouter.PUT("/:id", middlewares.AdminAuth(), controllers.UpdateMovie)
		movieRouter.DELETE("/:id", middlewares.AdminAuth(), controllers.DeleteMovie)
	}
}

func ReviewRoutes(router *gin.Engine) {

	reviewRouter := router.Group("/reviews")

	{
		reviewRouter.POST("/", middlewares.Auth(), controllers.CreateReview)
		reviewRouter.PUT("/:id", middlewares.Auth(), controllers.UpdateReview)
		reviewRouter.DELETE("/:id", middlewares.AdminAuth(), controllers.DeleteReview)
		reviewRouter.GET("/:id", middlewares.Auth(), controllers.GetReviewByMovieId)
	}
}
