package repository_test

import (
	"testing"

	"github.com/jivfur/profile-service/internal/model"
	"github.com/jivfur/profile-service/internal/repository"
	"github.com/jivfur/profile-service/repository"
	"github.com/stretchr/testify/assert"
)

// This would be replaced with your actual implementation or a mock

func TestCreateProfile_Success(t *testing.T) {
	mockRepo := &repository.MockProfileRepository{
		CreateFunc: func(p model.Profile) error {
			return nil // Simulate successful creation
		},
	}

	profile := model.Profile{
		ID:   "test-id",
		Name: "Updated User",
		// Update other fields
	}

	err := createProfile(mockRepo, profile)
	assert.NoError(t, err)

}

// func TestGetProfileByID(t *testing.T) {
// 	id := "test-id"
// 	profile, err := repo.GetByID(id)
// 	assert.NoError(t, err)
// 	assert.Equal(t, id, profile.ID)
// }

// func TestUpdateProfile(t *testing.T) {
// 	profile := model.Profile{
// 		ID:   "test-id",
// 		Name: "Updated User",
// 		// Update other fields
// 	}

// 	err := repo.Update(profile)
// 	assert.NoError(t, err)
// }

// func TestDeleteProfile(t *testing.T) {
// 	id := "test-id"
// 	err := repo.Delete(id)
// 	assert.NoError(t, err)
// }
