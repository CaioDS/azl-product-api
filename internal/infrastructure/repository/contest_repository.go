package repository

import (
	"azl-vote-api/internal/domain/entities"
	"database/sql"
)

type ContestRepository struct {
	DbConnection *sql.DB
}

func NewContestRepository(dbConnection *sql.DB) *ContestRepository {
	return &ContestRepository{
		DbConnection: dbConnection,
	}
}

func (contestRepository *ContestRepository) Create(entity entities.ContestEntity) error {
	_, error := contestRepository.DbConnection.Exec("INSERT INTO Contests(id, initial_date, final_date, active) VALUES (?,?,?,?)",
		entity.Id, entity.InitialDate.String(), entity.FinalDate.String(), entity.Active)

	if error != nil {
		return error
	}

	return nil
}
