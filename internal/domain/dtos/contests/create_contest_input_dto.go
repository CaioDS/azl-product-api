package contests

import "time"

type CreateContestInputDto struct {
	InitialDate time.Time `json:"initialDate"`
	FinalDate   time.Time `json:"finalDate"`
}
