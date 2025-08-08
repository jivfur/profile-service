package repository

import (
	"github.com/jivfur/profile-service/internal/model"
	"gorm.io/gorm"
)

type GormProfileRepository struct {
	db *gorm.DB
}

func NewGormProfileRepository(db *gorm.DB) *GormProfileRepository {
	return &GormProfileRepository{db: db}
}

func (r *GormProfileRepository) Create(profile model.Profile) error {
	return r.db.Create(&profile).Error
}

func (r *GormProfileRepository) GetByID(id string) (model.Profile, error) {
	var profile model.Profile
	err := r.db.First(&profile, "id = ?", id).Error
	return profile, err
}

func (r *GormProfileRepository) Update(profile model.Profile) error {
	return r.db.Save(&profile).Error
}

func (r *GormProfileRepository) Delete(id string) error {
	return r.db.Delete(&model.Profile{}, "id = ?", id).Error
}
