package spec

import "github.com/adm87/msgen/example/models"

// @Summary Example service
// @Description This is an example service specification, used to demonstrate the msgen tool.
type ExampleService interface {

	// @Summary Create a new user profile
	// @Description Creates a new user profile with the provided information.
	// @Version 1
	// @Param user body models.CreateUserRequest true
	// @Success 201 {object} models.UserProfile
	// @Failure 400 {object} models.ErrorResponse
	// @Failure 500 {object} models.ErrorResponse
	// @Router /users [post]
	CreateUserProfile(user models.CreateUserRequest) (models.UserProfile, error)

	// @Summary Delete user profile by ID
	// @Description Deletes the user profile associated with the given user ID.
	// @Version 1
	// @Param userID string path true
	// @Success 204 {string} string
	// @Failure 404 {object} models.ErrorResponse
	// @Failure 500 {object} models.ErrorResponse
	// @Router /users/{userID} [delete]
	DeleteUserProfile(userID string) error

	// @Summary Update user profile
	// @Description Updates the user profile information for a given user ID.
	// @Version 1
	// @Param userID string path true
	// @Param user body models.UserProfile true
	// @Success 200 {object} models.UserProfile
	// @Failure 400 {object} models.ErrorResponse
	// @Failure 404 {object} models.ErrorResponse
	// @Failure 500 {object} models.ErrorResponse
	// @Router /users/{userID} [put]
	UpdateUserProfile(userID string, user models.UserProfile) (models.UserProfile, error)

	// @Summary Get user profile by ID
	// @Description Retrieves the user profile information for a given user ID.
	// @Version 1
	// @Param userID string path true
	// @Success 200 {object} models.UserProfile
	// @Failure 404 {object} models.ErrorResponse
	// @Failure 500 {object} models.ErrorResponse
	// @Router /users/{userID} [get]
	GetUserProfile(userID string) (models.UserProfile, error)

	// @Summary Filter user profiles by age and country
	// @Description Retrieves a list of user profiles filtered by age and country.
	// @Version 1
	// @Param age int query false
	// @Param country string query false
	// @Success 200 {array} models.UserProfile
	// @Failure 500 {object} models.ErrorResponse
	// @Router /users/filter [get]
	FilterUserProfiles(age int, country string) ([]models.UserProfile, error)
}
