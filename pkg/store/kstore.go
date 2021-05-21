package store

import (
	"context"

	"time"

//	"entgo.io/ent/schema/edge"
	"github.com/ichandxyx/task/ent"
	"github.com/ichandxyx/task/ent/job"
	//"github.com/ichandxyx/task/ent/visit"
	//"github.com/ichandxyx/task/pkg/api"
)

type Visit struct {
	StoreID   string    `json:"store_id"`
	VisitTime time.Time `json:"visit_time"`
	ImageURLs []string  `json:"image_url"`
}

func (s *Storage) CreateJob(ctx context.Context) (*ent.Job, error) {

	return s.db.Job.Create().Save(ctx)

}

func (s *Storage) CreateVisits(ctx context.Context, visits []Visit, jobID int) ([]*ent.Visit, error) {

	bulk := make([]*ent.VisitCreate, len(visits))
	for i, vi := range visits {
		bulk[i] = s.db.Visit.Create().SetStoreID(vi.StoreID).SetVisitTime(vi.VisitTime).SetImageURLs(vi.ImageURLs).SetJobID(jobID)
	}

	return s.db.Visit.CreateBulk(bulk...).Save(ctx)
}

func (s *Storage) GetJob(ctx context.Context, jobID int) (*ent.Job, error) {
	return s.db.Job.Query().Where(job.ID(jobID)).WithVisits().First(ctx)
}

func (s *Storage) UpdateVisit(ctx context.Context, id int, peri int, er string) (*ent.Visit, error) {
	upd := s.db.Visit.UpdateOneID(id)
	if peri != 0 {
		upd.SetPerimeter(peri)
	}
	if er != "" {
		upd.SetNillableError(&er)
	}
	return upd.Save(ctx)
}

func (s *Storage) UpdateJob(ctx context.Context, id int, status string) (*ent.Job, error) {
	return s.db.Job.UpdateOneID(id).SetStatus(status).Save(ctx)
}
type Result struct{
	Store_id string `json:"store_id"`
	Area string      `json:"area"`
	Store_name string  `json:"store_name"`
	Data struct{
		Date time.Time	`json:"date"`
		Perimeter int	`json:"perimeter"`
	} `json:"data"`
}
func (s *Storage) GetVisits(ctx context.Context ,area string,sd time.Time,ed time.Time)([]Result){
	var res []Result
	for _,i:=range s.master{
		if i.Area==area{
			temp,_:= s.db.Job.Query().Where(job.CreatedAtGT(sd),job.CreatedAtLT(ed)).QueryVisits().Select(i.StoreID).All(ctx)
			Sname:=i.StoreName
			for _,j:=range temp{
				var tres Result
				tres.Area=area
				tres.Store_id=i.StoreID
				tres.Store_name=Sname
				tres.Data.Date=j.VisitTime
				tres.Data.Perimeter=j.Perimeter
				res = append(res,tres)

			}
		}
	}
	return res 


}

