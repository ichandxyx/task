package store

import (
	"context"
	
	"image"
	"os"
	"time"

	"github.com/ichandxyx/task/ent"
	//"github.com/ichandxyx/task/pkg/api"
)
type Visit struct{
	StoreID string `json:"store_id`
	VisitTime time.Time `json:"visit_time"`
	ImageURLs []string `json:"image_url"`

}

func (s *Storage) CreateJob(ctx context.Context) (*ent.Job,error){

	return s.db.Job.Create().Save(ctx)

}


func (s *Storage) CreateVisits(ctx context.Context, visits []Visit, jobID int) ([]*ent.Visit,error){
	
	bulk:=make([]*ent.VisitCreate,len(visits))
	for i, vi:=range visits{
		bulk[i]=s.db.Visit.Create().SetStoreID(vi.StoreID).SetVisitTime(vi.VisitTime).SetImageURLs(vi.ImageURLs).SetJobID(jobID)
	}

	return s.db.Visit.CreateBulk(bulk...).Save(ctx)
}


var p int
func (s *Storage) UpdatePerimeter(ctx context.Context,images []string){
	for _,x:=range images{
		width, height := getImageDimension(ctx,s,x)
		p+=(width+height)
	}
	s.db.Visit.Create().SetPerimeter(p).Save(ctx)
}


func getImageDimension(context context.Context,s *Storage ,imagePath string) (int, int) {
    file, err := os.Open(imagePath)
    if err != nil {
        s.db.Visit.Create().SetError("cant open URL").Save(context)
    }

    image, _, err := image.DecodeConfig(file)
    if err != nil {
        s.db.Visit.Create().SetError("cant decode").Save(context)

    }
    return image.Width, image.Height
}
type Response struct{
	StoreID string
}
func (s *Storage) GetStatus(ctx context.Context, jobID int)(*Response,error){
	RT,err:=s.db.Job.Query()
	return RT,err
}
type Result struct{
	Store_id string 
	area string
	Store_name string
	data []struct{
		Date time.Time
		Perimeter int
	}
}
func (s *Storage)GetVisits(ctx context.Context,qr  Query)(Result,error){
	var Re Result
	return Re,err

}
