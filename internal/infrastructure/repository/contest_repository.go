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

func (contestRepository *ContestRepository) FindAll() ([]*entities.ContestEntity, error) {
	rows, error := contestRepository.DbConnection.Query("SELECT id, initial_date, final_date, active FROM Contests")
	if error != nil {
		return nil, error
	}
	defer rows.Close()

	var contests []*entities.ContestEntity
	for rows.Next() {
		var contest entities.ContestEntity

		error = rows.Scan(&contest.Id, &contest.InitialDate, &contest.FinalDate, &contest.Active)
		if error != nil {
			return nil, error
		}

		contests = append(contests, &contest)
	}

	return contests, nil
}
