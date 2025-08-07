package repository

import "github.com/jivfur/profile-service/internal/model"

type MockProfileRepository struct {
	CreateFunc  func(profile model.Profile) error
	GetByIDFunc func(id string) (model.Profile, error)
	UpdateFunc  func(profile model.Profile) error
	DeleteFunc  func(id string) error
}

func (m *MockProfileRepository) Create(profile model.Profile) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(profile)
	}
	return nil
}

func (m *MockProfileRepository) GetByID(id string) (model.Profile, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(id)
	}
	return model.Profile{}, nil
}
func (m *MockProfileRepository) Update(profile model.Profile) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(profile)
	}
	return nil
}
func (m *MockProfileRepository) Delete(id string) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(id)
	}
	return nil
}
