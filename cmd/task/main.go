package main

import (
	"context"
	"net/http"
	"os"

	"entgo.io/ent/entc/integration/multischema/ent"
	"github.com/go-chi/chi/v5"
	"github.com/go-kit/kit/log"
	_ "github.com/mattn/go-sqlite3"
	"github.com/go-kit/kit/log/level"
	"github.com/ichandxyx/task/pkg/store"
	"github.com/ichandxyx/task/pkg/api"
)

func main() {

	logger :=log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger=log.With(logger,"ts",log.DefaultTimestampUTC,"caller",log.DefaultCaller)

	client,err:=ent.Open("sqllite3","file:"+"task.db"+"?cache=shared&mode=rwc&_fk=1")
	if err !=nil{
		level.Error(logger).Log("err",err)

	}
	level.Info(logger).Log("msg", "successfully connected to database")
	if err := client.Schema.Create(context.Background()); err != nil {
		level.Error(logger).Log("msg", "auto migration", "err", err)
	}

	st:=store.New(client)
	ap:=api.New(st, log.With(logger, "component", "api"))
	r:=chi.NewRouter()
	r.Route("/api",func(r chi.Router){
		ap.Register(r)
	})

	level.Info(logger).Log("msg", "starting web server", "addr", "0.0.0.0:8080")
	if err := http.ListenAndServe("0.0.0.0:8080", r); err != nil && err != http.ErrServerClosed {
		level.Error(logger).Log("msg", "http server listen", "err", err)
	}


}