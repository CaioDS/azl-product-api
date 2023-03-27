package contests

import (
	dtos "azl-vote-api/internal/domain/dtos/contests"
)

type CreateContestUsecase interface {
	Execute(input dtos.CreateContestInputDto) (*dtos.CreateContestOutputDto, error)
}
