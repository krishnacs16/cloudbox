package repository

import "services/iam/internal/models"

type AccessKeyRepository interface {
	Create(key *models.AccessKey) error

	GetByAccessKey(key string) (*models.AccessKey, error)

	ListByUser(userID int64) ([]models.AccessKey, error)

	Delete(id int64) error
}