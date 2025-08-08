package service

import (
	"errors"
	"net/mail"
	"time"

	"github.com/jivfur/profile-service/internal/model"
	"github.com/jivfur/profile-service/internal/repository"
)

type ProfileService struct {
	repo repository.ProfileRepository // Dependency injection for the repository
}

var Genders = map[string]bool{
	"MALE":        true,
	"FEMALE":      true,
	"TRANS MAN":   true,
	"TRANS WOMAN": true,
	"NON BINARY":  true,
	"QUEER":       true,
}

var SexualOrientation = map[string]bool{
	"GAY":        true,
	"STRAIGHT":   true,
	"BISEXUAL":   true,
	"ASEXUAL":    true,
	"DEMISEXUAL": true,
	"PANSEXUAL":  true,
	"QUEER":      true,
}

var SexualPosition = map[string]bool{
	"TOP":         true,
	"VERS TOP":    true,
	"VERSATILE":   true,
	"VERS BOTTOM": true,
	"BOTTOM":      true,
	"SIDE":        true,
}

func NewProfileService(repo repository.ProfileRepository) *ProfileService {
	return &ProfileService{repo: repo}
}

func (s *ProfileService) CreateProfile(profile model.Profile) error {

	if profile.Email == "" || !isValidEmail(profile.Email) {
		return errors.New("email is required")
	}

	if profile.PasswordHash == "" {
		return errors.New("password is required")
	}

	if len(profile.PasswordHash) != 60 {
		return errors.New("password is not the correct length")
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

	if _, exists := model.Genders[profile.Gender]; profile.Gender != "" && !exists {
		return errors.New("gender is not valid")
	}

	if _, exists := model.SexualOrientation[profile.SexualOrientation]; profile.SexualOrientation != "" && !exists {
		return errors.New("sexual orientation is not valid")
	}

	if _, exists := model.SexualPosition[profile.SexualPosition]; profile.SexualPosition != "" && !exists {
		return errors.New("sexual position is not valid")
	}

	// orchestrate calls to repo
	return s.repo.Create(profile)
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
