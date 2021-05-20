package store

import (
	"github.com/ichandxyx/task/ent"
)

type Storage struct {
	db *ent.Client
	master []Master
}
type Master struct{
	Area string`json:"area"`
	StoreName string `json:"store_name"`
	StoreID string `json:"store_id"`
}
func New(client *ent.Client ,mast []Master) *Storage {
	return &Storage{
		db: client,
		master :mast,
	}
}
