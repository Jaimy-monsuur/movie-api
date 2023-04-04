package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jaimy-monsuur/movie-api/src/Responses"
	"github.com/jaimy-monsuur/movie-api/src/Responses/exceptions"
	"github.com/jaimy-monsuur/movie-api/src/dtos"
	"github.com/jaimy-monsuur/movie-api/src/services"
)

// LoginUser godoc
// @Summary      login user with valid email and password combination
// @Description  login user
// @Tags         Auth
// @Security  BasicAuth
// @Accept       json
// @Produce      json
// @Param 		 data	body	dtos.LoginUserDto	true	"User Login Credentials JSON"
// @Success      200  {object}  dtos.SuccessResponseDto	"login successful"
// @Failure      400  {object}  dtos.FailedResponseDto	"request body validation errors"
// @Failure      401  {object}  dtos.FailedResponseDto	"invalid credentials"
// @Failure      500  {object}  dtos.FailedResponseDto	"unexpected internal server error"
// @Router       /auth/login [post]
func LoginUser(context *gin.Context) {

	// Validate Request Body
	body := dtos.LoginUserDto{}

	if err := context.BindJSON(&body); err != nil {
		exceptions.HandleValidationException(context, err)
		return
	}

	userExists, err := services.GetUserByEmail(body.Email)

	if err != nil {

		exceptions.HandleUnauthorizedException(context, "Invalid Credentials")
		return
	}

	if invalidPasswordError := userExists.ValidatePassword(body.Password); invalidPasswordError != nil {

		exceptions.HandleUnauthorizedException(context, "Invalid Credentials")
		return

	}

	accessToken, err := services.GenerateJwt(userExists.ID)
	if err != nil {

		exceptions.HandleInternalServerException(context)
		return
	}

	Responses.HandleOkResponse(context, "Login Successful", accessToken)
}
