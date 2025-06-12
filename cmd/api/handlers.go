package main

import "net/http"

type healthCheckResponse struct {
	Status string `json: "status"`
}

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	body := healthCheckResponse{Status: "OK"}
	err := app.writeJSON(w, http.StatusOK, body, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
