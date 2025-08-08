package service_test

import (
	"errors"
	"testing"
	"time"

	"github.com/jivfur/profile-service/internal/model"
	"github.com/jivfur/profile-service/internal/repository"
	"github.com/jivfur/profile-service/internal/service"
	"github.com/jivfur/profile-service/internal/service/testhelper"
	"github.com/stretchr/testify/assert"
)

// This would be replaced with your actual implementation or a mock

func TestCreateProfile_Success(t *testing.T) {
	mockRepo := &repository.MockProfileRepository{
		CreateFunc: func(p model.Profile) error {
			return nil // Simulate successful creation
		},
	}

	testCases := []struct {
		name     string
		profile  model.Profile
		expected error
	}{
		{name: "Valid Profile",
			profile:  *testhelper.NewFakeProfile(testhelper.WithCustom("DateOfBirth", time.Now().AddDate(-20, 0, 0))),
			expected: nil},
		{name: "Missing ID", profile: model.Profile{
			Name:         "Test User",
			Email:        "fake@email.com",
			PasswordHash: "hashedpassword",
		}, expected: errors.New("ID is required")},
		{name: "Missing Email", profile: model.Profile{
			ID:           "test-id",
			Name:         "Test User",
			PasswordHash: "hashedpassword",
		}, expected: errors.New("email is required")},
		{name: "Missing Password", profile: model.Profile{
			ID:    "test-id",
			Name:  "Test User",
			Email: "fake@email.com",
		}, expected: errors.New("password is required")},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Logf("Test case: %s, Profile: %+v", tc.name, tc.profile)
			svc := service.NewProfileService(mockRepo)
			err := svc.CreateProfile(tc.profile)
			if tc.expected == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expected.Error())
			}
		})
	}
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
