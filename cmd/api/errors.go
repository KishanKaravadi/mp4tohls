package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Logger struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func NewLogger(errorOutput, infoOutput io.Writer) *Logger {
	return &Logger{
		errorLog: log.New(errorOutput, "[ERROR]:\t", log.Ldate|log.Ltime|log.Lshortfile),
		infoLog:  log.New(infoOutput, "[INFO]:\t", log.Ldate|log.Ltime),
	}
}

func (l *Logger) Info(msg string) {
	l.infoLog.Println(msg)
}

func (l *Logger) Error(msg error) {
	l.errorLog.Println(msg)
}

func (l *Logger) Fatal(msg error) {
	l.errorLog.Fatalln(msg)
}

func (app *application) logError(r *http.Request, err error) {
	reqInfo, marshallErr := json.MarshalIndent(map[string]string{
		"request_method": r.Method,
		"request_url":    r.URL.String(),
	}, "", "    ")
	if marshallErr != nil {
		panic(marshallErr)
	}
	app.logger.errorLog.Println(err, string(reqInfo))

}

type errorResponse struct {
	Error string `json:"error"`
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message string) {
	err := app.writeJSON(w, status, errorResponse{Error: message}, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	message := "the server encountered a problem and could not process yopur request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}
