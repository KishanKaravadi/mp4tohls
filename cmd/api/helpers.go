package main

import (
	"encoding/json"
	"maps"
	"net/http"
)

func (app *application) writeJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	jsonData = append(jsonData, '\n')
	maps.Copy(w.Header(), headers)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}
