package dependencyinjections

import (
	usecaseImpl "azl-vote-api/internal/application/usecase/contests"
	repository "azl-vote-api/internal/domain/repository"
	usecase "azl-vote-api/internal/domain/usecase/contests"
)

func ProvideCreateContestUsecase(repository repository.ContestRepository) usecase.CreateContestUsecase {
	return usecaseImpl.NewCreateContestUsecase(repository)
}
