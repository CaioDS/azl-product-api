package models

import "time"

type Contest struct {
	Id          string
	InitialDate time.Time
	FinalDate   time.Time
	Active      bool
}
