 
package api

import (
	"context"
	

	"github.com/go-chi/chi/v5"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/ichandxyx/task/ent"

	//"github.com/ichandxyx/task/ent/visit"
	//"github.com/ichandxyx/task/pkg/img"
	"github.com/ichandxyx/task/pkg/store"
	"image"
    _ "image/jpeg"
    "net/http"
)

type API struct {
	store  *store.Storage
	logger log.Logger
}

func New(s *store.Storage, logger log.Logger) *API {
	return &API{
		store:  s,
		logger: logger,
	}
}
func (a *API) Register(r chi.Router) {
	a.registerRoutes(r)
}

func (a *API) calculatePerimeter(ctx context.Context, jobid int, visits []*ent.Visit) {
	// perimeters:=make( map[string]int)
	// errs:=make(map[string]string)
	status := "completed"
	for _, vi := range visits {
		var errFlag bool
		var p int
		//a.store.UpdatePerimeter(ctx,&(vi.ImageURLs))
		for _, i := range vi.ImageURLs {
			resp, err := http.Get(i)
			if err!=nil{
				errFlag = true
			}
			defer resp.Body.Close()
			config, _, err:=image.DecodeConfig(resp.Body)
			//config, err := img.ImageConfigFromURL(i)
			if err != nil {
				// perimeters[vi.StoreID]=-1
				// errs[vi.StoreID]="image download failed"
				errFlag = true
					//ahndle errer
					level.Error(a.logger).Log("err", err)
				break
			}
			p += 2 * (config.Height + config.Width)
		}
		if errFlag {
			a.store.UpdateVisit(ctx, vi.ID, 0, "image download failed")
			status = "failed"
			continue
		}
		a.store.UpdateVisit(ctx, vi.ID, p, "")

	}
	a.store.UpdateJob(ctx, jobid, status)

	//a=visit[0].perimeter
	//a.store.UpdatePerimeter(ctx,visit)

}
