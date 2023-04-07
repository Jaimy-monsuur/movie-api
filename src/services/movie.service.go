package services

import (
	"errors"
	"github.com/jaimy-monsuur/movie-api/src/config"
	"github.com/jaimy-monsuur/movie-api/src/dtos"
	"github.com/jaimy-monsuur/movie-api/src/interfaces"
	"github.com/jaimy-monsuur/movie-api/src/models"
)

func CreateMovie(movie dtos.CreateMovie) (*models.Movie, *interfaces.ServiceError) {
	//check if movie already exists
	var movieExists models.Movie

	if err := config.DB.First(&movieExists, "title = ?", movie.Title).Error; err == nil {
		movieExistsError := &interfaces.ServiceError{
			Error:      errors.New("Movie with title: " + movie.Title + " already exists"),
			StatusCode: 409,
		}
		return nil, movieExistsError
	}

	newMovie := models.Movie{
		Title:    movie.Title,
		Year:     movie.Year,
		Director: movie.Director,
		Actors:   movie.Actors,
		Plot:     movie.Plot,
		Language: movie.Language,
		Length:   movie.Length,
		Url:      movie.Url,
	}

	result := config.DB.Create(&newMovie)

	if result.Error != nil {
		movieCreateError := &interfaces.ServiceError{
			Error:      result.Error,
			StatusCode: 400,
		}
		return nil, movieCreateError
	}

	return &newMovie, nil
}

func GetAllMovies() ([]*models.Movie, error) {
	var allMovies []*models.Movie

	err := config.DB.Preload("Reviews").Preload("Reviews.User").Find(&allMovies).Error

	if err != nil {
		return nil, err
	}

	return allMovies, nil
}

func GetMovieById(ID string) (*models.Movie, error) {
	var movie models.Movie

	if err := config.DB.Preload("Reviews").Preload("Reviews.User").First(&movie, "id = ?", ID).Error; err != nil {

		return nil, err
	}

	return &movie, nil
}

func UpdateMovie(id string, movie dtos.UpdateMovie) (*models.Movie, error) {
	var movieToUpdate models.Movie

	err := config.DB.First(&movieToUpdate, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	movieToUpdate.Title = movie.Title
	movieToUpdate.Year = movie.Year
	movieToUpdate.Director = movie.Director
	movieToUpdate.Actors = movie.Actors
	movieToUpdate.Plot = movie.Plot
	movieToUpdate.Language = movie.Language
	movieToUpdate.Length = movie.Length
	movieToUpdate.Url = movie.Url

	err = config.DB.Save(&movieToUpdate).Error

	if err != nil {
		return nil, err
	}

	return &movieToUpdate, nil
}

func DeleteMovie(id string) error {
	var movieToDelete models.Movie

	err := config.DB.Delete(&movieToDelete, "id = ?", id)

	if err.Error != nil {
		return err.Error
	}

	return nil
}
