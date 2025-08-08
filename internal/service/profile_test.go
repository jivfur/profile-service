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
		{
			name:     "Valid Profile",
			profile:  *testhelper.NewFakeProfile(),
			expected: nil,
		},
		{
			name:     "Missing Email",
			profile:  *testhelper.NewFakeProfile(testhelper.WithNo("Email")),
			expected: errors.New("email is required"),
		},
		{
			name:     "Invalid Email",
			profile:  *testhelper.NewFakeProfile(testhelper.WithCustom("Email", "inv@lid-email@domain.com")),
			expected: errors.New("email is required"),
		},
		{
			name:     "Missing Password",
			profile:  *testhelper.NewFakeProfile(testhelper.WithNo("PasswordHash")),
			expected: errors.New("password is required"),
		},
		{
			name:     "Password too short",
			profile:  *testhelper.NewFakeProfile(testhelper.WithCustom("PasswordHash", "short")),
			expected: errors.New("password is not the correct length"),
		},
		{
			name:     "Underage Profile",
			profile:  *testhelper.NewFakeProfile(testhelper.WithCustom("DateOfBirth", time.Now().AddDate(-15, 0, 0))),
			expected: errors.New("date of birth must be at least 18 years ago"),
		},
		{
			name:     "Underage Profile",
			profile:  *testhelper.NewFakeProfile(testhelper.WithCustom("DateOfBirth", time.Now().AddDate(+15, 0, 0))),
			expected: errors.New("date of birth must be at least 18 years ago"),
		},
		{
			name:     "Fake Gender",
			profile:  *testhelper.NewFakeProfile(testhelper.WithCustom("Gender", "FAKE GENDER")),
			expected: errors.New("gender is not valid"),
		},
		{
			name:     "No Gender",
			profile:  *testhelper.NewFakeProfile(testhelper.WithNo("Gender")),
			expected: nil, // No error
		},
		{
			name:     "Fake Sexual Orientation",
			profile:  *testhelper.NewFakeProfile(testhelper.WithCustom("SexualOrientation", "FAKE GENDER")),
			expected: errors.New("sexual orientation is not valid"),
		},
		{
			name:     "No Sexual Orientation",
			profile:  *testhelper.NewFakeProfile(testhelper.WithNo("SexualOrientation")),
			expected: nil, // No error
		},
		{
			name:     "Fake Sexual Position",
			profile:  *testhelper.NewFakeProfile(testhelper.WithCustom("SexualPosition", "FAKE GENDER")),
			expected: errors.New("sexual position is not valid"),
		},
		{
			name:     "No Sexual Position",
			profile:  *testhelper.NewFakeProfile(testhelper.WithNo("SexualPosition")),
			expected: nil, // No error
		},
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
