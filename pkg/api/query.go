package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ichandxyx/task/pkg/store"
)

func (a *API) handleStatus(w http.ResponseWriter,r *http.Request){
	var input struct{
		JobID int `json:"jobid"`
	}
	err:=json.NewDecoder(r.Body).Decode(&input)
	if err!=nil{
		//handle error
	}
	status,err:=a.store.GetStatus(r.Context(),input.JobID)
	if err!=nil{
		//handle error
	}
	json.NewEncoder(w).Encode(status)
}

func (a *API) handleSubmit(w http.ResponseWriter,r *http.Request){
	var input struct {
		Count int `json:"count"`
		visits []store.Visit `json:"visits"`
	}
	err:=json.NewDecoder(r.Body).Decode(&input)
	if err!=nil {
		//ahndle errer
	}
	job,err:=a.store.CreateJob(r.Context())
	visits,err:=a.store.CreateVisits(r.Context(), input.visits, job.ID)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"job_id":job.ID})

	a.calculatePerimeter(r.Context(),visits)
}
type Query struct{
	Area int `json:"area"`
	StoreID int `json:"storeid"`
	StartDate time.Time `json:"stdate"`
	EndDate time.Time `json:"endate"`
}
func (a *API) handleVisits(w http.ResponseWriter,r *http.Request){
	var qr Query
	err:=json.NewDecoder(r.Body).Decode(&qr)
	if err!=nil{
		//handle error
	}
	res,err:=a.store.GetVisits(r.Context(),qr)
		if err!=nil{
			w.WriteHeader(http.StatusBadRequest)
		}
		json.NewEncoder(w).Encode(&res)

}


