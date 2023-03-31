package contests

import (
	dtos "azl-vote-api/internal/domain/dtos/contests"
)

type ListContestsUsecase interface {
	Execute() ([]*dtos.ListContestsOutputDto, error)
}
