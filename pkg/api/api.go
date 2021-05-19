package api

import (
	"log"
	"github.com/go-chi/chi/v5"
	"github.com/ichandxyx/task/pkg/store"
)
type API struct{
	store *store.Storage
	logger log.Logger
}

func New( s *store.Storage, logger log.Logger)*API {
	return &API{
		store:s,
		logger:logger,
	}
}
func (a *API) Register(r chi.Router) {
	a.registerRoutes(r)
}