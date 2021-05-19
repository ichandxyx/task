package api

import (
	"github.com/go-chi/chi/v5"
)

func (a *API) registerRoutes(r chi.Router){
	r.Get("/status",a.handleStatus)
	r.Post("/submit",a.handleSubmit)
	r.Get("/visits",a.handleVisits)
}