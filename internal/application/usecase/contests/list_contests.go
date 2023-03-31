package contests

import (
	dtos "azl-vote-api/internal/domain/dtos/contests"
	"azl-vote-api/internal/domain/repository"
)

type ListContestsUsecase struct {
	ContestRepository repository.ContestRepository
}

func NewListContestsUsecase(contestRepository repository.ContestRepository) *ListContestsUsecase {
	return &ListContestsUsecase{
		ContestRepository: contestRepository,
	}
}

func (usecase *ListContestsUsecase) Execute() ([]*dtos.ListContestsOutputDto, error) {
	contests, error := usecase.ContestRepository.FindAll()
	if error != nil {
		return nil, error
	}

	var contestsOutput []*dtos.ListContestsOutputDto
	for _, contest := range contests {
		contestsOutput = append(contestsOutput, &dtos.ListContestsOutputDto{
			Id:          contest.Id,
			InitialDate: contest.InitialDate,
			FinalDate:   contest.FinalDate,
			Active:      contest.Active,
		})
	}

	return contestsOutput, nil
}
