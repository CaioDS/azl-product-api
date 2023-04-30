package repository

import (
	"azl-vote-api/internal/domain/entities"
)

type ContestRepository interface {
	FindAll() ([]*entities.ContestEntity, error)
	Create(*entities.ContestEntity) error
	DeleteById(*entities.ContestEntity) error
}
