package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jaimy-monsuur/movie-api/src/Responses"
	"github.com/jaimy-monsuur/movie-api/src/Responses/exceptions"
	"github.com/jaimy-monsuur/movie-api/src/dtos"
	"github.com/jaimy-monsuur/movie-api/src/services"
)

// CreateMovie godoc
// @Summary Create a movie
// @Description Create a movie
// @Tags Movie
// @Accept json
// @Produce json
// @Param data body dtos.CreateMovie true "New Movie Details JSON"
// @Success 201 {object} dtos.SuccessResponseDto{data=models.Movie} "movie created successfully"
// @Failure 400 {object} dtos.FailedResponseDto "request body validation error"
// @Failure 409 {object} dtos.FailedResponseDto "movie with supplied title already exists"
// @Failure 500 {object} dtos.FailedResponseDto "unexpected internal server error"
// @Router /movies [post]
func CreateMovie(context *gin.Context) {
	//validate request body
	body := dtos.CreateMovie{}

	if err := context.BindJSON(&body); err != nil {

		exceptions.HandleValidationException(context, err)
		return
	}

	newMovie, err := services.CreateMovie(body)

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

	Responses.HandleCreatedResponse(context, "Movie Created", newMovie)
}

// GetAllMovies godoc
// @Summary Get all movies
// @Description Get all movies
// @Tags Movie
// @Security JWT
// @Accept json
// @Produce json
// @Success 200 {object} dtos.SuccessResponseDto{data=[]models.Movie} "all movies returned"
// @Failure 500 {object} dtos.FailedResponseDto "unexpected internal server error"
// @Router /movies [get]
func GetAllMovies(context *gin.Context) {
	movies, err := services.GetAllMovies()

	if err != nil {
		exceptions.HandleInternalServerException(context)
		return
	}

	Responses.HandleOkResponse(context, "Movies returned", movies)
}

// UpdateMovie godoc
// @Summary Update a movie
// @Description Update a movie
// @Tags Movie
// @Security JWT
// @Accept json
// @Produce json
// @Param id path int true "Movie ID"
// @Param data body dtos.UpdateMovie true "Update Movie Details JSON"
// @Success 200 {object} dtos.SuccessResponseDto{data=models.Movie} "movie updated successfully"
// @Failure 400 {object} dtos.FailedResponseDto "request body validation error"
// @Failure 404 {object} dtos.FailedResponseDto "movie not found"
// @Failure 409 {object} dtos.FailedResponseDto "movie with supplied title already exists"
// @Failure 500 {object} dtos.FailedResponseDto "unexpected internal server error"
// @Router /movies/{id} [put]
func UpdateMovie(context *gin.Context) {
	//validate Request Params
	id := dtos.EntityID{}

	if err := context.ShouldBindUri(&id); err != nil {
		exceptions.HandleValidationException(context, err)
		return
	}

	//validate request body
	body := dtos.UpdateMovie{}

	if err := context.BindJSON(&body); err != nil {

		exceptions.HandleValidationException(context, err)
		return
	}

	movie, err := services.UpdateMovie(id.ID, body)

	if err != nil {
		exceptions.HandleNotFoundException(context, err)
		return

	}

	Responses.HandleOkResponse(context, "Movie Updated", movie)
}

// DeleteMovie godoc
// @Summary Delete a movie
// @Description Delete a movie
// @Tags Movie
// @Security JWT
// @Accept json
// @Produce json
// @Param id path int true "Movie ID"
// @Success 200 {object} dtos.SuccessResponseDto{data=models.Movie} "movie deleted successfully"
// @Failure 404 {object} dtos.FailedResponseDto "movie not found"
// @Failure 500 {object} dtos.FailedResponseDto "unexpected internal server error"
// @Router /movies/{id} [delete]
func DeleteMovie(context *gin.Context) {
	//validate Request Params
	id := dtos.EntityID{}
	if err := context.ShouldBindUri(&id); err != nil {
		exceptions.HandleValidationException(context, err)
		return
	}

	err := services.DeleteMovie(id.ID)

	if err != nil {
		exceptions.HandleNotFoundException(context, err)
		return

	}

	Responses.HandleOkResponse(context, "Movie Deleted", nil)
}

// GetMovieByID godoc
// @Summary Get a movie
// @Description Get a movie
// @Tags Movie
// @Security JWT
// @Accept json
// @Produce json
// @Param id path int true "Movie ID"
// @Success 200 {object} dtos.SuccessResponseDto{data=models.Movie} "movie returned"
// @Failure 404 {object} dtos.FailedResponseDto "movie not found"
// @Failure 500 {object} dtos.FailedResponseDto "unexpected internal server error"
// @Router /movies/{id} [get]
func GetMovieByID(context *gin.Context) {
	//validate Request Params
	params := dtos.EntityID{}

	if err := context.BindUri(&params); err != nil {
		exceptions.HandleValidationException(context, err)
		return
	}

	movie, err := services.GetMovieById(params.ID)

	if err != nil {
		exceptions.HandleNotFoundException(context, err)
		return

	}

	Responses.HandleOkResponse(context, "Movie returned", movie)
}
