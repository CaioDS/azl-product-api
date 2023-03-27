package dependencyinjections

import (
	"azl-vote-api/internal/infrastructure/repository"
	"database/sql"
)

func ProvideContestRepository(dbConnection *sql.DB) *repository.ContestRepository {
	return repository.NewContestRepository(dbConnection)
}
