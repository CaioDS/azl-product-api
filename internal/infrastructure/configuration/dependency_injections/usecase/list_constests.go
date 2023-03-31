package dependencyinjections

import (
	usecaseImpl "azl-vote-api/internal/application/usecase/contests"
	repository "azl-vote-api/internal/domain/repository"
	usecase "azl-vote-api/internal/domain/usecase/contests"
)

func ProvideListContestsUsecase(repository repository.ContestRepository) usecase.ListContestsUsecase {
	return usecaseImpl.NewListContestsUsecase(repository)
}
