package middlewares

import (
	"errors"
	"github.com/jaimy-monsuur/movie-api/src/Responses/exceptions"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jaimy-monsuur/movie-api/src/services"
)

func Auth() gin.HandlerFunc {

	return func(context *gin.Context) {
		bearerToken := context.GetHeader("Authorization")
		if bearerToken == "" {
			exceptions.HandleBadRequestException(context, errors.New("bearer token is required"))
			return
		}

		accessToken := strings.Split(bearerToken, "Bearer ")[1]
		err := services.ValidateToken(accessToken)
		if err != nil {

			exceptions.HandleUnauthorizedException(context, "Unauthorized")
			return
		}
		context.Next()
	}
}

func AdminAuth() gin.HandlerFunc {

	return func(context *gin.Context) {
		bearerToken := context.GetHeader("Authorization")
		if bearerToken == "" {
			exceptions.HandleBadRequestException(context, errors.New("bearer token is required"))
			return
		}

		accessToken := strings.Split(bearerToken, "Bearer ")[1]
		err := services.ValidateAdminToken(accessToken)
		if err != nil {

			exceptions.HandleUnauthorizedException(context, "Unauthorized")
			return
		}
		context.Next()
	}
}
