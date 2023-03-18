package repository

import (
	"azl-product-api/internal/domain/entities"
	"database/sql"
)

type VoteRepository struct {
	DbConnection *sql.DB
}

func NewVoteRepository(dbConnection *sql.DB) *VoteRepository {
	return &VoteRepository{DbConnection: dbConnection}
}

func (voteRepository *VoteRepository) Create(vote *entities.VoteEntity) error {
	_, error := voteRepository.DbConnection.Exec("INSERT INTO VOTES (id) values (?)", vote.Id)

	if error != nil {
		return error
	}
	return nil
}
