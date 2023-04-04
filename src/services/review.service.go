package services

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jaimy-monsuur/movie-api/src/config"
	"github.com/jaimy-monsuur/movie-api/src/dtos"
	"github.com/jaimy-monsuur/movie-api/src/interfaces"
	"github.com/jaimy-monsuur/movie-api/src/models"
)

func CreateReview(context *gin.Context, review dtos.CreateReviewDto) (*models.Review, *interfaces.ServiceError) {
	//get user
	user, err := GetUserByID(review.UserID)

	if err != nil {
		userNotFoundError := &interfaces.ServiceError{
			Error:      err,
			StatusCode: 404,
		}
		return nil, userNotFoundError
	}

	//get movie
	movie, err := GetMovieById(review.MovieID)

	if err != nil {
		movieNotFoundError := &interfaces.ServiceError{
			Error:      err,
			StatusCode: 404,
		}
		return nil, movieNotFoundError
	}

	newReview := models.Review{
		UserID:  uuid.MustParse(user.ID),
		MovieID: movie.ID,
		Content: review.Review,
		Rating:  float64(review.Rating),
	}

	//check if user from token is the same as the user in the request body
	err = CheckUser(context, user.ID)

	if err != nil {
		userUnauthorizedError := &interfaces.ServiceError{
			Error:      err,
			StatusCode: 401,
		}
		return nil, userUnauthorizedError
	}

	result := config.DB.Create(&newReview)

	if result.Error != nil {
		reviewCreateError := &interfaces.ServiceError{
			Error:      result.Error,
			StatusCode: 400,
		}
		return nil, reviewCreateError
	}

	return &newReview, nil
}

func GetReviewsByMovieId(ID string) ([]*models.Review, error) {
	var reviews []*models.Review

	err := config.DB.Where("movie_id = ?", ID).Find(&reviews).Error

	if err != nil {
		return nil, err
	}

	return reviews, nil

}

func UpdateReview(context *gin.Context, ID string, review dtos.UpdateReviewDto) (*models.Review, *interfaces.ServiceError) {
	var reviewToUpdate models.Review

	err := config.DB.First(&reviewToUpdate, "ID = ?", ID).Error

	if err != nil {
		reviewNotFoundError := &interfaces.ServiceError{
			Error:      err,
			StatusCode: 404,
		}
		return nil, reviewNotFoundError
	}

	//check if user from token is the same as the user in the request body
	err = CheckUser(context, reviewToUpdate.UserID.String())

	if err != nil {
		userUnauthorizedError := &interfaces.ServiceError{
			Error:      err,
			StatusCode: 401,
		}
		return nil, userUnauthorizedError
	}

	reviewToUpdate.Content = review.Review
	reviewToUpdate.Rating = float64(review.Rating)

	result := config.DB.Save(&reviewToUpdate)

	if result.Error != nil {
		reviewUpdateError := &interfaces.ServiceError{
			Error:      result.Error,
			StatusCode: 400,
		}
		return nil, reviewUpdateError
	}

	return &reviewToUpdate, nil
}

func DeleteReview(ID string) *interfaces.ServiceError {
	var reviewToDelete models.Review

	err := config.DB.First(&reviewToDelete, "ID = ?", ID).Error

	if err != nil {
		reviewNotFoundError := &interfaces.ServiceError{
			Error:      err,
			StatusCode: 404,
		}
		return reviewNotFoundError
	}

	result := config.DB.Delete(&reviewToDelete)

	if result.Error != nil {
		reviewDeleteError := &interfaces.ServiceError{
			Error:      result.Error,
			StatusCode: 400,
		}
		return reviewDeleteError
	}

	return nil
}
