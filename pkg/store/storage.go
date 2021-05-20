package store

import (
	"github.com/ichandxyx/task/ent"

)

type Storage struct {
	db *ent.Client
}

func New(client *ent.Client) *Storage{
	return &Storage{
		db:client,
	}
}