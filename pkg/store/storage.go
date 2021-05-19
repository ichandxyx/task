package store

import (
	"entgo.io/ent/entc/integration/multischema/ent"
)

type Storage struct {
	db *ent.Client
}

func New(client *ent.Client) *Storage{
	return &Storage{
		db:client,
	}
}