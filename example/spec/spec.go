package spec

import "github.com/adm87/msgen/example/models"

type ExampleService interface {
	// @Summary Get user profile by ID
	// @Param userID string path true
	// @Router /users/{userID} [get]
	GetUserProfile(userID string) (models.UserProfile, error)

	// @Summary Filter user profiles by age and country
	// @Param age int query false
	// @Param country string query false
	// @Router /users/filter [get]
	FilterUserProfiles(age int, country string) ([]models.UserProfile, error)
}
