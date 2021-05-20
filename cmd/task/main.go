package main

import (
	"context"
	"encoding/csv"
	"net/http"
	"os"

	"github.com/ichandxyx/task/ent"
	"github.com/ichandxyx/task/pkg/api"
	"github.com/ichandxyx/task/pkg/store"

	"github.com/go-chi/chi/v5"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	client, err := ent.Open("sqlite3", "file:task.db?cache=shared&mode=rwc&_fk=1")
	if err != nil {
		level.Error(logger).Log("err", err)

	}
	level.Info(logger).Log("msg", "successfully connected to database")
	if err := client.Schema.Create(context.Background()); err != nil {
		level.Error(logger).Log("msg", "auto migration", "err", err)
	}
	storesFile,err:=os.Open("./stores.csv")
	if err!=nil{
		level.Error(logger).Log("msg", "opening stores.csv file", "err", err)

	}
	cr:=csv.NewReader(storesFile)
	storeMaster,err:=cr.ReadAll()
	if err!=nil{
		level.Error(logger).Log("msg", "parsing stores.csv", "err", err)
	}
	sm:=make([]store.Master,0,len(storeMaster))
	for _,i:=range storeMaster{
		sm=append(sm, store.Master{Area: i[0],StoreName: i[1],StoreID: i[2]})
	}
	st := store.New(client,sm)
	ap := api.New(st, log.With(logger, "component", "api"))
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		ap.Register(r)
	})

	level.Info(logger).Log("msg", "starting web server", "addr", "0.0.0.0:8080")
	if err := http.ListenAndServe("0.0.0.0:8080", r); err != nil && err != http.ErrServerClosed {
		level.Error(logger).Log("msg", "http server listen", "err", err)
	}

}
