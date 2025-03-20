package service

import (
	"errors"
	"wedding/models"
	"wedding/repository"
)

type AllService struct {
	repo *repository.Repository
}

func NewAllService(repo *repository.Repository) *AllService {
	return &AllService{repo: repo}
}

func (s *AllService) GetAll() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *AllService) GetById(user *models.User, id string) error {
	err := s.repo.GetById(user, id)
	if err != nil {
		return errors.New("user not found")
	}
	return nil
} 

func (s *AllService) GetByName(user *models.User, name string) error {
    // Melakukan pencarian berdasarkan nama, bukan ID
    err := s.repo.GetByName(user, name)
    if err != nil {
        return errors.New("user not found")
    }
    return nil
}


func (s *AllService) AddData(User *models.User) error {
	if User.Name == "" {
		return errors.New("name is required")
	}
	return s.repo.AddData(User)
}

func (s *AllService) EditData(User *models.User) error {
	return s.repo.EditData(User)
}

func (s *AllService) DeleteData(id uint) error {
	return s.repo.DeleteData(id)
}
