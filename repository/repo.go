package repository

import (
	"fmt"
	"log"
	"wedding/config"
	"wedding/models"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) GetAll() ([]models.User, error) {
    var products []models.User
    result := r.DB.Find(&products)
    return products, result.Error
}

func (r *Repository) GetById(User *models.User, id string) error {
    log.Println("Menerima ID:", id) // Log untuk memeriksa ID yang diterima

    if id == "" {
        return fmt.Errorf("ID tidak valid")
    }

    result := r.DB.First(User, "id = ?", id)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func (r *Repository) GetByName(User *models.User, name string) error {
	result := r.DB.First(User, "name = ?", name)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Repository) AddData(User *models.User) error {
	if config.DB == nil {
		return fmt.Errorf("database no connceted")
	}
	return r.DB.Create(User).Error
}

func (r *Repository) EditData(User *models.User) error {
	return r.DB.Save(User).Error
}

func (r *Repository) DeleteData(id uint) error {
	return r.DB.Delete(&models.User{}, id).Error
}
