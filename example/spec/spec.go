package spec

import "github.com/adm87/msgen/example/models"

// @Summary Example service
// @Description This is an example service specification, used to demonstrate the msgen tool.
type ExampleService interface {

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
