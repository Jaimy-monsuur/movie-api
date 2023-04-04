package main

import (
	"github.com/jaimy-monsuur/movie-api/src/config"
	"github.com/jaimy-monsuur/movie-api/src/models"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
}

func main() {
	config.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	config.DB.AutoMigrate(&models.Movie{}, &models.Review{}, &models.User{})
}
