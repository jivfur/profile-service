package service

import (
	"errors"
	"time"

	"github.com/jivfur/profile-service/internal/model"
	"github.com/jivfur/profile-service/internal/repository"
)

type ProfileService struct {
	repo repository.ProfileRepository // Dependency injection for the repository
}

func NewProfileService(repo repository.ProfileRepository) *ProfileService {
	return &ProfileService{repo: repo}
}

func (s *ProfileService) CreateProfile(profile model.Profile) error {
	// validate profile data
	if profile.ID == "" {
		return errors.New("ID is required")
	}

	if profile.Email == "" {
		return errors.New("email is required")
	}

	if profile.PasswordHash == "" {
		return errors.New("password is required")
	}

	if profile.DateOfBirth == nil {
		return errors.New("date of birth is required")
	}

	today := time.Now()
	// Calculate the date 18 years ago from today
	eighteenYearsAgo := today.AddDate(-18, 0, 0)
	if profile.DateOfBirth.After(eighteenYearsAgo) {
		return errors.New("date of birth must be at least 18 years ago")
	}

	// orchestrate calls to repo
	return s.repo.Create(profile)
}
