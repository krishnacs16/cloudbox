package repository

import "services/iam/internal/models"

type UserRepository interface {

	Create(user *models.User) error
	GetByID(id int64) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	List() ([]models.User, error)
	Delete(id int64) error
	
}