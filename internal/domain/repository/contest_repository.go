package repository

import "azl-vote-api/internal/domain/entities"

type ContestRepository interface {
	Create(entities.ContestEntity) error
	FindAll() ([]*entities.ContestEntity, error)
}
