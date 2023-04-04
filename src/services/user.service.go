package services

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jaimy-monsuur/movie-api/src/config"
	"github.com/jaimy-monsuur/movie-api/src/dtos"
	"github.com/jaimy-monsuur/movie-api/src/interfaces"
	"github.com/jaimy-monsuur/movie-api/src/models"
)

func CreateUser(createUserDto *dtos.CreateUserDto) (*models.User, *interfaces.ServiceError) {

	// Check if Another User with specified email already exists
	var userExists models.User

	if err := config.DB.First(&userExists, "email = ?", createUserDto.Email).Error; err == nil {

		userExistsError := &interfaces.ServiceError{
			Error:      errors.New("User with email: " + createUserDto.Email + " already exists"),
			StatusCode: 409,
		}

		return nil, userExistsError
	}

	newUser := models.User{
		Email:     createUserDto.Email,
		Password:  createUserDto.Password,
		FirstName: createUserDto.FirstName,
		LastName:  createUserDto.LastName,
	}

	result := config.DB.Create(&newUser)

	if result.Error != nil {

		userCreateError := &interfaces.ServiceError{
			Error:      result.Error,
			StatusCode: 400,
		}
		return nil, userCreateError
	}

	newUser.Password = ""

	return &newUser, nil
}

func GetAllUsers() ([]*dtos.UserDto, error) {

	var allUsers []*models.User

	err := config.DB.Select("id", "email", "first_name", "last_name", "last_login", "created_at", "updated_at").Find(&allUsers).Error

	if err != nil {
		return nil, err
	}

	var returnUsers []*dtos.UserDto

	for _, user := range allUsers {
		var userDto dtos.UserDto
		userDto.ID = user.ID.String()
		userDto.Email = user.Email
		userDto.FirstName = user.FirstName
		userDto.LastName = user.LastName
		userDto.Role = user.Role
		returnUsers = append(returnUsers, &userDto)
	}

	return returnUsers, nil
}

func GetUserByEmail(email string) (*models.User, error) {

	var user models.User

	if err := config.DB.First(&user, "email = ?", email).Error; err != nil {

		return nil, err
	}

	return &user, nil
}

func GetUserByID(userID string) (*dtos.UserDto, error) {

	var user models.User

	if err := config.DB.First(&user, "id = ?", userID).Error; err != nil {

		return nil, err
	}

	var userDto dtos.UserDto
	userDto.ID = user.ID.String()
	userDto.Email = user.Email
	userDto.FirstName = user.FirstName
	userDto.LastName = user.LastName
	userDto.Role = user.Role

	return &userDto, nil
}

func UpdateUser(context *gin.Context, userID string, updateUserDto *dtos.UpdateUserDto) (*models.User, *interfaces.ServiceError) {

	// Check if User with supplied ID exists
	var user models.User
	result := config.DB.First(&user, "id = ?", userID)

	if result.Error != nil {

		//create service error
		userNotFoundError := &interfaces.ServiceError{
			Error:      result.Error,
			StatusCode: 404,
		}
		return nil, userNotFoundError
	}

	//check if user from token is the same as the user from the request
	err := CheckUser(context, user.ID.String())

	if err != nil {
		unauthorized := &interfaces.ServiceError{
			Error:      err,
			StatusCode: 401,
		}
		return nil, unauthorized
	}

	config.DB.Model(&user).Updates(models.User{
		FirstName: updateUserDto.FirstName,
		LastName:  updateUserDto.LastName,
		Email:     updateUserDto.Email,
	})

	user.Password = ""

	return &user, nil

}

func DeleteUser(userID string) error {

	result := config.DB.Delete(&models.User{}, "id = ?", userID)

	if result.Error != nil {
		return result.Error
	}

	return nil

}
