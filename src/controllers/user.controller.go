package controllers

import (
	"github.com/jaimy-monsuur/movie-api/src/Responses"
	"github.com/jaimy-monsuur/movie-api/src/Responses/exceptions"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaimy-monsuur/movie-api/src/dtos"
	"github.com/jaimy-monsuur/movie-api/src/services"
)

// CreateUser godoc
// @Summary      registers a new user
// @Description  create user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param 		 data	body	dtos.CreateUserDto	true	"New User Details JSON"
// @Success      201  {object}  dtos.SuccessResponseDto{data=models.User}	"user created successfully"
// @Failure      400  {object}  dtos.FailedResponseDto "request body validation error"
// @Failure      409  {object}  dtos.FailedResponseDto "another user with supplied email exists"
// @Failure      500  {object}  dtos.FailedResponseDto "unexpected internal server error"
// @Router       /users [post]
func CreateUser(context *gin.Context) {

	// Validate Request Body
	body := dtos.CreateUserDto{}

	if err := context.BindJSON(&body); err != nil {

		exceptions.HandleValidationException(context, err)
		return
	}

	newUser, err := services.CreateUser(&body)

	if err != nil {

		switch statusCode := err.StatusCode; statusCode {
		case 400:
			exceptions.HandleBadRequestException(context, err.Error)
			return
		case 409:
			exceptions.HandleConflictException(context, err.Error.Error())
			return
		}
	}

	Responses.HandleCreatedResponse(context, "User Registered", newUser)
}

// GetAllUsers godoc
// @Summary      returns all users
// @Description  get all users
// @Tags         User
// @Security 	JWT
// @Accept       json
// @Produce      json
// @success 200 {object} dtos.SuccessResponseDto{data=[]models.User}	"all users returned"
// @Failure      400  {object}  dtos.FailedResponseDto	"token not passed with request"
// @Failure      401  {object}  dtos.FailedResponseDto	"invalid/expired token"
// @Failure      500  {object}  dtos.FailedResponseDto	"unexpected internal server error"
// @Router       /users [get]
func GetAllUsers(context *gin.Context) {

	allUsers, err := services.GetAllUsers()

	if err != nil {
		exceptions.HandleInternalServerException(context)
		return
	}

	Responses.HandleOkResponse(context, "All Users", allUsers)
}

// GetUserByID godoc
// @Summary      returns a user by its 16 caharcter uuid
// @Description  get user by ID
// @Tags         User
// @Security 	JWT
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID(UUID)"
// @success 200 {object} dtos.SuccessResponseDto{data=models.User} "desc"
// @Failure      400  {object}  dtos.FailedResponseDto	"request param validation error or token not passed with request"
// @Failure      401  {object}  dtos.FailedResponseDto	"invalid/expired token"
// @Failure      404  {object}  dtos.FailedResponseDto	"user with the specified ID not found"
// @Failure      500  {object}  dtos.FailedResponseDto	"unexpected internal server error"
// @Router       /users/{id} [get]
func GetUserByID(context *gin.Context) {

	// Validate Request Params
	params := dtos.EntityID{}

	if err := context.BindUri(&params); err != nil {
		exceptions.HandleValidationException(context, err)
		return
	}

	user, err := services.GetUserByID(params.ID)

	if err != nil {
		exceptions.HandleNotFoundException(context, err)
		return
	}

	Responses.HandleOkResponse(context, "User with ID: "+params.ID, user)
}

// UpdateUser godoc
// @Summary      updates a user
// @Description  update user
// @Tags         User
// @Security 	JWT
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID(UUID)"
// @Param 		 data	body	dtos.UpdateUserDto	true	"User Details JSON"
// @success 200 {object} dtos.SuccessResponseDto{data=models.User}	"user updated successfully"
// @Failure      400  {object}  dtos.FailedResponseDto	"request body/param validation error or token not passed with request"
// @Failure      401  {object}  dtos.FailedResponseDto	"invalid/expired token"
// @Failure      404  {object}  dtos.FailedResponseDto	"user with specified ID not found"
// @Failure      500  {object}  dtos.FailedResponseDto	"unexpected internal server error"
// @Router       /users/{id} [put]
func UpdateUser(context *gin.Context) {

	// Validate Request Params
	params := dtos.EntityID{}
	if err := context.BindUri(&params); err != nil {
		exceptions.HandleValidationException(context, err)
		return
	}

	// Validate Request Body
	body := dtos.UpdateUserDto{}
	if err := context.BindJSON(&body); err != nil {
		exceptions.HandleValidationException(context, err)
		return
	}

	user, err := services.UpdateUser(context, params.ID, &body)

	if err != nil {
		switch statusCode := err.StatusCode; statusCode {
		case 401:
			exceptions.HandleUnauthorizedException(context, "Unauthorized")
			return
		case 404:
			exceptions.HandleNotFoundException(context, err.Error)
			return
		}
	}

	Responses.HandleOkResponse(context, "User Updated", user)

}

// DeleteUser godoc
// @Summary      deletes a user
// @Description  delete user
// @Tags         User
// @Security 	JWT
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID(UUID)"
// @success 200 {object} dtos.SuccessResponseDto	"user deleted suuceesfully"
// @Failure      400  {object}  dtos.FailedResponseDto	"request param validation error or token not passed with request"
// @Failure      401  {object}  dtos.FailedResponseDto	"invalid/expired token"
// @Failure      500  {object}  dtos.FailedResponseDto	"unexpected internal server error"
// @Router       /users/{id} [delete]
func DeleteUser(context *gin.Context) {

	// Validate Request Params
	params := dtos.EntityID{}

	if err := context.BindUri(&params); err != nil {
		exceptions.HandleValidationException(context, err)
		return
	}

	if err := services.DeleteUser(params.ID); err != nil {
		exceptions.HandleBadRequestException(context, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"statusText": "success",
		"statusCode": 200,
		"message":    "User Deleted",
	})
}
