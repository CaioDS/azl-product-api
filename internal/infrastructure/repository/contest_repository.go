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

func (contestRepository *ContestRepository) Create(entity *entities.ContestEntity) error {
	dateLayout := "2006-01-02 15:04:05 -07:00"
	parsedInitialDate := entity.InitialDate.Format(dateLayout)
	parsedFinalDate := entity.FinalDate.Format(dateLayout)

	statement, error := contestRepository.DbConnection.Prepare("INSERT INTO Contests(id, initial_date, final_date, active) VALUES (@p1, @p2, @p3, @p4)")
	if error != nil {
		return error
	}

	_, error = statement.Exec(entity.Id, parsedInitialDate, parsedFinalDate, entity.Active)
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
