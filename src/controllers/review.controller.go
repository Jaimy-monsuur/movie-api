package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jaimy-monsuur/movie-api/src/Responses"
	"github.com/jaimy-monsuur/movie-api/src/Responses/exceptions"
	"github.com/jaimy-monsuur/movie-api/src/dtos"
	"github.com/jaimy-monsuur/movie-api/src/services"
)

// CreateReview godoc
// @Summary Create a review
// @Description Create a review
// @Tags Review
// @Accept json
// @Produce json
// @Param data body dtos.CreateReviewDto true "New Review Details JSON"
// @Success 201 {object} dtos.SuccessResponseDto{data=models.Review} "review created successfully"
// @Failure 400 {object} dtos.FailedResponseDto "request body validation error"
// @Failure 409 {object} dtos.FailedResponseDto "review with supplied title already exists"
// @Failure 500 {object} dtos.FailedResponseDto "unexpected internal server error"
// @Router /reviews [post]
func CreateReview(context *gin.Context) {
	//validate request body
	body := dtos.CreateReviewDto{}

	if err := context.BindJSON(&body); err != nil {

		exceptions.HandleValidationException(context, err)
		return
	}

	newReview, err := services.CreateReview(context, body)

	if err != nil {

		switch statusCode := err.StatusCode; statusCode {
		case 400:
			exceptions.HandleBadRequestException(context, err.Error)
			return
		case 409:
			exceptions.HandleConflictException(context, err.Error.Error())
			return
		case 404:
			exceptions.HandleNotFoundException(context, err.Error)
			return
		case 401:
			exceptions.HandleUnauthorizedException(context, "Unauthorized")
			return
		}
	}

	Responses.HandleCreatedResponse(context, "Review Created", newReview)
}

// GetReviewByMovieId godoc
// @Summary Get a review by movie id
// @Description Get a review by movie id
// @Tags Review
// @Accept json
// @Produce json
// @Param id path int true "Movie ID"
// @Success 200 {object} dtos.SuccessResponseDto{data=models.Review} "review returned"
// @Failure 404 {object} dtos.FailedResponseDto "review not found"
// @Failure 500 {object} dtos.FailedResponseDto "unexpected internal server error"
// @Router /reviews/{id} [get]
func GetReviewByMovieId(context *gin.Context) {
	id := dtos.EntityID{}

	if err := context.BindUri(&id); err != nil {
		exceptions.HandleValidationException(context, err)
		return
	}

	review, err := services.GetReviewsByMovieId(id.ID)

	if err != nil {
		exceptions.HandleNotFoundException(context, err)
		return
	}

	Responses.HandleOkResponse(context, "Review returned", review)
}

// UpdateReview godoc
// @Summary Update a review
// @Description Update a review
// @Tags Review
// @Accept json
// @Produce json
// @Param id path int true "Review ID"
// @Param data body dtos.UpdateReviewDto true "Update Review Details JSON"
// @Success 200 {object} dtos.SuccessResponseDto{data=models.Review} "review updated successfully"
// @Failure 400 {object} dtos.FailedResponseDto "request body validation error"
// @Failure 404 {object} dtos.FailedResponseDto "review not found"
// @Failure 500 {object} dtos.FailedResponseDto "unexpected internal server error"
// @Router /reviews/{id} [put]
func UpdateReview(context *gin.Context) {
	id := dtos.EntityID{}

	if err := context.BindUri(&id); err != nil {
		exceptions.HandleValidationException(context, err)
		return
	}

	//validate request body
	body := dtos.UpdateReviewDto{}

	if err := context.BindJSON(&body); err != nil {

		exceptions.HandleValidationException(context, err)
		return
	}

	updatedReview, err := services.UpdateReview(context, id.ID, body)

	if err != nil {

		switch statusCode := err.StatusCode; statusCode {
		case 400:
			exceptions.HandleBadRequestException(context, err.Error)
			return
		case 404:
			exceptions.HandleNotFoundException(context, err.Error)
			return
		case 401:
			exceptions.HandleUnauthorizedException(context, "Unauthorized")
			return
		}
	}

	Responses.HandleOkResponse(context, "Review Updated", updatedReview)
}

// DeleteReview godoc
// @Summary Delete a review
// @Description Delete a review
// @Tags Review
// @Accept json
// @Produce json
// @Param id path int true "Review ID"
// @Success 200 {object} dtos.SuccessResponseDto{data=models.Review} "review deleted successfully"
// @Failure 404 {object} dtos.FailedResponseDto "review not found"
// @Failure 500 {object} dtos.FailedResponseDto "unexpected internal server error"
// @Router /reviews/{id} [delete]
func DeleteReview(context *gin.Context) {
	id := dtos.EntityID{}

	if err := context.BindUri(&id); err != nil {
		exceptions.HandleValidationException(context, err)
		return
	}

	err := services.DeleteReview(id.ID)

	if err != nil {
		exceptions.HandleNotFoundException(context, err.Error)
		return

	}

	Responses.HandleOkResponse(context, "Review Deleted", nil)
}
