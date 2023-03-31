package contests

import (
	dtos "azl-vote-api/internal/domain/dtos/contests"
	"azl-vote-api/internal/domain/entities"
	"azl-vote-api/internal/domain/repository"
)

type CreateContestUsecase struct {
	ContestRepository repository.ContestRepository
}

func NewCreateContestUsecase(contestRepository repository.ContestRepository) *CreateContestUsecase {
	return &CreateContestUsecase{
		ContestRepository: contestRepository,
	}
}

func (usecase *CreateContestUsecase) Execute(input dtos.CreateContestInputDto) (*dtos.CreateContestOutputDto, error) {
	contest := entities.CreateContestEntity(input.InitialDate, input.FinalDate)

	error := usecase.ContestRepository.Create(contest)
	if error != nil {
		return nil, error
	}

	return &dtos.CreateContestOutputDto{
		Id:          contest.Id,
		InitialDate: contest.InitialDate,
		FinalDate:   contest.FinalDate,
		Active:      contest.Active,
	}, nil
}
