package entities

import "github.com/docker/distribution/uuid"

type VoteEntity struct {
	Id string
}

func CreateVote(name string, description string, price float64) *VoteEntity {
	return &VoteEntity{
		Id: string(uuid.Generate().String()),
	}
}
