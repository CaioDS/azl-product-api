package entities

import (
	"time"

	"github.com/docker/distribution/uuid"
)

type ContestEntity struct {
	Id          string
	InitialDate time.Time
	FinalDate   time.Time
	Active      bool
}

func CreateContestEntity(initialDate time.Time, finalDate time.Time) *ContestEntity {
	return &ContestEntity{
		Id:          uuid.Generate().String(),
		InitialDate: initialDate,
		FinalDate:   finalDate,
		Active:      true,
	}
}
