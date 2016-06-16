package router

import (
	"github.com/gorilla/mux"
	"github.com/zjwdmlmx/ssensor/controllers"
)

var R *mux.Router

func init() {
	R = mux.NewRouter()

	R.HandleFunc("/sensor/data", controllers.SensorDataJsonHandler).
		HeadersRegexp("Content-Type", "application/json(\\;.*){0,1}").
		Methods("POST")

	R.HandleFunc("/sensor/data", controllers.SensorDataFormHandler).
		HeadersRegexp("Content-Type", "multipart/form-data(\\;.*){0,1}").
		Methods("POST")

	R.HandleFunc("/history/data", controllers.HistoryJsonHandler).
		HeadersRegexp("Content-Type", "application/json(\\;.*){0,1}").
		Methods("POST")

}
