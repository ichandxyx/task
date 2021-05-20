package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"github.com/go-kit/kit/log/level"
	"github.com/ichandxyx/task/pkg/store"
)

func (a *API) handleStatus(w http.ResponseWriter, r *http.Request) {


	
	jobIdRaw:=r.URL.Query().Get("jobid")
	jobId,err:=strconv.Atoi(jobIdRaw)
	if err != nil {
		//handle error
		level.Error(a.logger).Log("err", err)
		return
	}
	job, err := a.store.GetJob(r.Context(), jobId)
	if err != nil {
		//handle error
		level.Error(a.logger).Log("err", err)
		return
	}
	type resper struct{
			StoreID string `json:"store_id"`
			Error string `json:"error"`
		}
	var resp struct{
		Status string `json:"status"`
		JobID int `json:"job_id"`
		Error []resper `json:"error,omitempty"`
	}
	resp.Status=job.Status
	resp.JobID=job.ID
	resp.Error =make([]resper, 0)
	for _,vi:=range job.Edges.Visits{
		if vi.Error!=nil{
			resp.Error=append(resp.Error, resper{
				StoreID:vi.StoreID,
				Error:*vi.Error,

			})
		}
	}
	

	json.NewEncoder(w).Encode(resp)
}

func (a *API) handleSubmit(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Count  int           `json:"count"`
		Visits []store.Visit `json:"visits"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		//ahndle errer
		level.Error(a.logger).Log("err", err)
		return
	}
	job, err := a.store.CreateJob(r.Context())
	visits, err := a.store.CreateVisits(r.Context(), input.Visits, job.ID)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"job_id": job.ID})

	go a.calculatePerimeter(context.Background(),job.ID, visits)
}


type Query struct {
	Area      string       `json:"area"`
	StoreID   string       `json:"storeid"`
	StartDate time.Time `json:"startdate"`
	EndDate   time.Time `json:"enddate"`
}

func (a *API) handleVisits(w http.ResponseWriter, r *http.Request) {
	var qr Query
	qr.Area=r.URL.Query().Get("area")
	qr.StoreID=r.URL.Query().Get("storeid")
	
	//t, err := time.Parse(time.RFC3339, str)
	qr.StartDate,_=time.Parse(time.RFC3339,r.URL.Query().Get("startdate"))
	qr.EndDate,_=time.Parse(time.RFC3339,r.URL.Query().Get("enddate"))
	
	res, err := a.store.GetVisits(r.Context(), qr.Area,qr.StartDate,qr.EndDate)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(&res)

}
