package api

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/go-kit/kit/log"
	"github.com/ichandxyx/task/ent"
	//"github.com/ichandxyx/task/ent/visit"
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
func(a *API) calculatePerimeter(ctx context.Context,visit []*ent.Visit){
	
	for _,vi:=range visit{
		//a.store.UpdatePerimeter(ctx,&(vi.ImageURLs))
		a.store.UpdatePerimeter(ctx,vi.ImageURLs)

	}
	//a=visit[0].perimeter
		//a.store.UpdatePerimeter(ctx,visit)
	

}
