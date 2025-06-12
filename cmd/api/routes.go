package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/v1/healthcheck", app.healthCheckHandler)

	// r.Post("/v1/upload", app.uploadHandler)
	// r.Post("/v1/process", app.processHandler)
	// r.Post("/v1/process/stop", app.stopProcessHandler)

	// r.Get("/v1/jobs", app.getJobHandler)
	// r.Get("/v1/videos", app.getVideoHandler)
	// r.Get("/v1/videos/{video_id}/download/hls", app.downloadHLSHandler)

	return r

}
