package contests

import "time"

type ListContestsOutputDto struct {
	Id          string
	InitialDate time.Time
	FinalDate   time.Time
	Active      bool
}