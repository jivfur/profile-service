package repository

import "github.com/jivfur/profile-service/internal/model"

type ProfileRepository interface {
	Create(profile model.Profile) error
	GetByID(id string) (model.Profile, error)
	Update(profile model.Profile) error
	Delete(id string) error
	// Add more methods as your app needs them
}
