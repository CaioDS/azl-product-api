package contests

import "time"

type CreateContestOutputDto struct {
	Id          string
	InitialDate time.Time
	FinalDate   time.Time
	Active      bool
}
